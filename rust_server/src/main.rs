use serde::{Deserialize, Serialize};
use sqlx::{postgres::PgPoolOptions, Pool, Postgres};
use axum::{
    extract::{Path, State}, http::StatusCode, response::IntoResponse, routing::get, Json, Router
};

#[derive(Clone)]
struct AppState {
    db: Pool<Postgres>,
}

#[derive(Debug, Deserialize, Serialize, sqlx::FromRow)]
pub struct UserModel {
    pub id: i32,
    pub name: String,
    pub email: String,
    pub created_at: chrono::NaiveDateTime,
}

async fn current_user(State(data): State<AppState>, Path(id): Path<i32>) -> impl IntoResponse {
    let query_result = sqlx::query_as!(UserModel, "SELECT * FROM users WHERE id = $1", id)
        .fetch_one(&data.db)
        .await;
    match query_result {
        Ok(user) => {
            let user_response = serde_json::json!({"status": "success","data": serde_json::json!({
                "user": user
            })});

            return (StatusCode::OK, Json(user_response));
        }
        Err(sqlx::Error::RowNotFound) => {
            return (StatusCode::NOT_FOUND, Json(serde_json::json!({"status": "fail","message": format!("User with ID: {} not found", id)})));
        }
        Err(e) => {
            return (StatusCode::INTERNAL_SERVER_ERROR, Json(serde_json::json!({"status": "error","message": format!("{:?}", e)})));
        }
    };
}

#[tokio::main]
async fn main() {
    dotenvy::dotenv().expect(".env file not found");
    let database_url = std::env::var("DATABASE_URL").expect("DATABASE_URL must be set");
    let pool = match PgPoolOptions::new()
        .max_connections(10)
        .connect(&database_url)
        .await
    {
        Ok(pool) => {
            println!("✅Connection to the database is successful!");
            pool
        }
        Err(err) => {
            println!("🔥 Failed to connect to the database: {:?}", err);
            std::process::exit(1);
        }
    };
    let app = Router::new()
        .route("/users/:id", get(current_user))
        .with_state(AppState {db: pool});
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3001")
        .await
        .unwrap();
    axum::serve(listener, app).await.unwrap();
}
