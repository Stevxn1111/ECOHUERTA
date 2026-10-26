import os
import psycopg2
import psycopg2.extras
from flask import Flask, render_template, abort

app = Flask(__name__,
            template_folder="../front/templates",
            static_folder="../front/static")

def consultar(sql, params=(), uno=False):
    conn = psycopg2.connect(
        host=os.getenv("DB_HOST", "localhost"),
        user=os.getenv("DB_USER", "postgres"),
        password=os.getenv("DB_PASSWORD", "123"),
        database=os.getenv("DB_NAME", "ECOHUERTA"),
    )
    try:
        with conn.cursor(cursor_factory=psycopg2.extras.RealDictCursor) as cur:
            cur.execute(sql, params)
            return cur.fetchone() if uno else cur.fetchall()
    finally:
        conn.close()

@app.route("/")
def mapa():
    return render_template("mapa.html")

@app.route("/departamento/<slug>")
def departamento(slug):
    dep = consultar("SELECT * FROM departamentos WHERE slug = %s", (slug,), uno=True)
    if not dep:
        abort(404)
    plantas = consultar(
        "SELECT id_planta, nombre, imagen FROM plantas WHERE id_departamento = %s",
        (dep["id_departamento"],),
    )
    return render_template("departamento.html", dep=dep, plantas=plantas)

@app.route("/planta/<int:id_planta>")
def planta(id_planta):
    p = consultar("SELECT * FROM plantas WHERE id_planta = %s", (id_planta,), uno=True)
    if not p:
        abort(404)
    recursos = consultar("SELECT * FROM recursos WHERE id_planta = %s", (id_planta,))
    return render_template("planta.html", p=p, recursos=recursos)

if __name__ == "__main__":
    app.run(debug=True)