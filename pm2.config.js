module.exports = {
    apps: [
        {
            name: "node_server",
            script: "./node_server/index.js",
            env_production: {
                NODE_ENV: "production"
            },
        },
        {
            name: 'rust_server',
            script: './rust_server/target/release/rust_server',
            exec_interpreter: 'none',
            exec_mode: 'fork_mode'
        },
        {
            name: 'go_server',
            script: './go_server/main',
            exec_interpreter: 'none',
            exec_mode: 'fork_mode'
        }
    ]
}