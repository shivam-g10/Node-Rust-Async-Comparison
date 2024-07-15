import express from 'express';
import pg from 'pg';
import dotenv from 'dotenv';

dotenv.config();

const app = express();
const db = new pg.Pool({
    connectionString: process.env.DATABASE_URL,
    min: 1,
    max: 10,
})


app.get('/users/:id', async (req, res) => {
    try {
        const result = await db.query('select * from users where id = $1', [req.params.id]);
        if (result.rows.length) {
            res.json({ 'status': 'success', 'data': { 'user': result.rows[0] } })
        } else {
            res.json({ 'status': 'fail', 'message': `User with ID: ${req.params.id} not found` })
        }
    } catch (error) {
        res.status(500).json({ 'status': 'error', 'message': error.message })
    }

})

app.listen(3000, "0.0.0.0", () => {
    console.info('App listening on port 3000')
})
