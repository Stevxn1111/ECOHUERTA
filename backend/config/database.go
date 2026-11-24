package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = os.Getenv("DB_PASSWORD")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" {
		user = "eco_user"
	}
	if pass == "" {
		pass = "tu_contraseña_segura"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3306"
	}
	if name == "" {
		name = "ecohuerta"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", user, pass, host, port, name)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir conexión con la base de datos: %v", err)
	}

	maxIntentos := 15
	for i := 1; i <= maxIntentos; i++ {
		if err = DB.Ping(); err == nil {
			break
		}
		log.Printf("Esperando a MariaDB (intento %d/%d): %v", i, maxIntentos, err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("No se pudo conectar a MariaDB tras %d intentos: %v", maxIntentos, err)
	}

	log.Println("Conexión exitosa a MariaDB")
	createTables()
	seedDepartamentos()
	seedTiposYPlantas()
}

func createTables() {
	usersQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL,
		usuario VARCHAR(50) NOT NULL UNIQUE,
		clave VARCHAR(255) NOT NULL,
		rol ENUM('estudiante', 'docente', 'admin') NOT NULL
	);`

	departamentosQuery := `
	CREATE TABLE IF NOT EXISTS departamentos (
		id_departamento INT AUTO_INCREMENT PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL,
		slug VARCHAR(50) NOT NULL UNIQUE
	);`

	tiposPlantaQuery := `
	CREATE TABLE IF NOT EXISTS tipos_planta (
		id_tipo INT AUTO_INCREMENT PRIMARY KEY,
		nombre VARCHAR(100) NOT NULL UNIQUE
	);`

	plantasQuery := `
	CREATE TABLE IF NOT EXISTS plantas (
		id_planta INT AUTO_INCREMENT PRIMARY KEY,
		nombre VARCHAR(120) NOT NULL,
		descripcion TEXT NOT NULL,
		caracteristicas TEXT NOT NULL,
		cuidados TEXT NOT NULL,
		beneficios TEXT NOT NULL,
		id_tipo INT NOT NULL,
		id_departamento INT NOT NULL,
		imagen VARCHAR(255) DEFAULT '',
		CONSTRAINT fk_planta_tipo FOREIGN KEY (id_tipo) REFERENCES tipos_planta(id_tipo),
		CONSTRAINT fk_planta_departamento FOREIGN KEY (id_departamento) REFERENCES departamentos(id_departamento)
	);`

	if _, err := DB.Exec(usersQuery); err != nil {
		log.Fatalf("Error al crear la tabla users: %v", err)
	}
	if _, err := DB.Exec(departamentosQuery); err != nil {
		log.Fatalf("Error al crear la tabla departamentos: %v", err)
	}
	if _, err := DB.Exec(tiposPlantaQuery); err != nil {
		log.Fatalf("Error al crear la tabla tipos_planta: %v", err)
	}
	if _, err := DB.Exec(plantasQuery); err != nil {
		log.Fatalf("Error al crear la tabla plantas: %v", err)
	}

	log.Println("Tablas 'users', 'departamentos', 'tipos_planta' y 'plantas' verificadas/creadas correctamente.")
}

func seedDepartamentos() {
	var count int
	if err := DB.QueryRow("SELECT COUNT(*) FROM departamentos").Scan(&count); err != nil {
		log.Printf("Aviso al consultar departamentos: %v", err)
		return
	}
	if count > 0 {
		return
	}

	departamentos := []struct {
		nombre string
		slug   string
	}{
		{"Amazonas", "amazonas"},
		{"Antioquia", "antioquia"},
		{"Arauca", "arauca"},
		{"Atlántico", "atlantico"},
		{"Bogotá D.C.", "bogota_dc"},
		{"Bolívar", "bolivar"},
		{"Boyacá", "boyaca"},
		{"Caldas", "caldas"},
		{"Caquetá", "caqueta"},
		{"Casanare", "casanare"},
		{"Cauca", "cauca"},
		{"Cesar", "cesar"},
		{"Chocó", "choco"},
		{"Córdoba", "cordoba"},
		{"Cundinamarca", "cundinamarca"},
		{"Guainía", "guainia"},
		{"Guaviare", "guaviare"},
		{"Huila", "huila"},
		{"La Guajira", "la_guajira"},
		{"Magdalena", "magdalena"},
		{"Meta", "meta"},
		{"Nariño", "narino"},
		{"Norte de Santander", "norte_de_santander"},
		{"Putumayo", "putumayo"},
		{"Quindío", "quindio"},
		{"Risaralda", "risaralda"},
		{"San Andrés y Providencia", "san_andres"},
		{"Santander", "santander"},
		{"Sucre", "sucre"},
		{"Tolima", "tolima"},
		{"Valle del Cauca", "valle_del_cauca"},
		{"Vaupés", "vaupes"},
		{"Vichada", "vichada"},
	}

	stmt, err := DB.Prepare("INSERT INTO departamentos (nombre, slug) VALUES (?, ?)")
	if err != nil {
		log.Printf("Error preparando inserción de departamentos: %v", err)
		return
	}
	defer stmt.Close()

	for _, d := range departamentos {
		if _, err := stmt.Exec(d.nombre, d.slug); err != nil {
			log.Printf("Error insertando departamento %s: %v", d.nombre, err)
		}
	}
	log.Println("Departamentos precargados correctamente.")
}

func seedTiposYPlantas() {
	tipos := []string{
		"Páramo y Alta Montaña",
		"Medicinal",
		"Ornamental",
		"Frutal y Alimenticia",
		"Bosque Tropical",
	}

	for _, nombreTipo := range tipos {
		_, err := DB.Exec("INSERT IGNORE INTO tipos_planta (nombre) VALUES (?)", nombreTipo)
		if err != nil {
			log.Printf("Error insertando tipo de planta %s: %v", nombreTipo, err)
		}
	}

	var count int
	if err := DB.QueryRow("SELECT COUNT(*) FROM plantas").Scan(&count); err != nil || count > 0 {
		return
	}

	var idParamo, idMedicinal, idOrnamental, idFrutal, idTropical int
	_ = DB.QueryRow("SELECT id_tipo FROM tipos_planta WHERE nombre = ?", "Páramo y Alta Montaña").Scan(&idParamo)
	_ = DB.QueryRow("SELECT id_tipo FROM tipos_planta WHERE nombre = ?", "Medicinal").Scan(&idMedicinal)
	_ = DB.QueryRow("SELECT id_tipo FROM tipos_planta WHERE nombre = ?", "Ornamental").Scan(&idOrnamental)
	_ = DB.QueryRow("SELECT id_tipo FROM tipos_planta WHERE nombre = ?", "Frutal y Alimenticia").Scan(&idFrutal)
	_ = DB.QueryRow("SELECT id_tipo FROM tipos_planta WHERE nombre = ?", "Bosque Tropical").Scan(&idTropical)

	getDeptoID := func(slug string) int {
		var id int
		_ = DB.QueryRow("SELECT id_departamento FROM departamentos WHERE slug = ?", slug).Scan(&id)
		return id
	}

	idBoyaca := getDeptoID("boyaca")
	idCundinamarca := getDeptoID("cundinamarca")
	idAntioquia := getDeptoID("antioquia")
	idAmazonas := getDeptoID("amazonas")
	idValle := getDeptoID("valle_del_cauca")

	type plantaSeed struct {
		nombre, desc, carac, cuidados, beneficios, imagen string
		idTipo, idDepto                                 int
	}

	semillas := []plantaSeed{
		{
			nombre: "Frailejón (Espeletia)",
			desc: "Planta emblemática de los páramos andinos con hojas densamente cubiertas de vellosidad plateada.",
			carac: "Crecimiento lento (unos pocos centímetros al año), roseta de hojas lanudas y tallo leñoso.",
			cuidados: "Requiere clima frío de páramo, suelos húmedos con buen drenaje y exposición solar de altura; no trasplantar de su hábitat.",
			beneficios: "Capta y regula el agua de la niebla y protege los suelos de alta montaña.",
			idTipo: idParamo, idDepto: idBoyaca, imagen: "",
		},
		{
			nombre: "Uchuva",
			desc: "Arbusto andino que produce pequeños frutos dorados encerrados en un cáliz papiráceo.",
			carac: "Tallos ramificados, fruto dulce-ácido rico en vitaminas y cáscara protectora natural.",
			cuidados: "Prefiere climas templados a fríos, suelos fértiles bien drenados y tutorado para sostener sus ramas.",
			beneficios: "Alta fuente de vitamina C, antioxidantes y aporte nutricional.",
			idTipo: idFrutal, idDepto: idBoyaca, imagen: "",
		},
		{
			nombre: "Caléndula",
			desc: "Hierba aromática de flores amarillo-anaranjadas muy usada en huertas educativas y botiquín natural.",
			carac: "Hojas alargadas suaves y capítulos florales brillantes que abren con la luz solar.",
			cuidados: "Sol directo o semisombra, riego moderado sin encharcar y suelo ligero.",
			beneficios: "Uso tradicional calmante en piel y atrae polinizadores benéficos.",
			idTipo: idMedicinal, idDepto: idCundinamarca, imagen: "",
		},
		{
			nombre: "Orquídea Cattleya",
			desc: "Una de las especies ornamentales más representativas de Colombia por el tamaño y color de sus flores.",
			carac: "Pseudobulbos engrosados, hojas carnosas y flores vistosas de aroma suave.",
			cuidados: "Luz filtrada brillante, buena ventilación, sustrato aireado para epífitas y riego controlado.",
			beneficios: "Alto valor ornamental y biodiversidad en bosques húmedos.",
			idTipo: idOrnamental, idDepto: idAntioquia, imagen: "",
		},
		{
			nombre: "Victoria Regia",
			desc: "Planta acuática amazónica famosa por sus enormes hojas circulares con borde levantado.",
			carac: "Hojas flotantes que pueden superar un metro de diámetro y flores que cambian de color.",
			cuidados: "Aguas cálidas y tranquilas, abundante luz solar y sustrato rico en materia orgánica.",
			beneficios: "Refugio para fauna acuática y equilibrio de ecosistemas de humedal.",
			idTipo: idTropical, idDepto: idAmazonas, imagen: "",
		},
		{
			nombre: "Chontaduro",
			desc: "Palma tropical de racimos con frutos rojizos o anaranjados de pulpa harinosa y nutritiva.",
			carac: "Tallo anillado con espinas y frutos ricos en grasas saludables y betacarotenos.",
			cuidados: "Clima cálido y húmedo, plena exposición solar y suelos profundos con buena humedad.",
			beneficios: "Alto aporte energético y nutricional en la gastronomía regional.",
			idTipo: idFrutal, idDepto: idValle, imagen: "",
		},
	}

	for _, p := range semillas {
		if p.idDepto == 0 || p.idTipo == 0 {
			continue
		}
		_, err := DB.Exec(`
			INSERT INTO plantas (nombre, descripcion, caracteristicas, cuidados, beneficios, id_tipo, id_departamento, imagen)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
			p.nombre, p.desc, p.carac, p.cuidados, p.beneficios, p.idTipo, p.idDepto, p.imagen,
		)
		if err != nil {
			log.Printf("Error insertando planta semilla %s: %v", p.nombre, err)
		}
	}
}