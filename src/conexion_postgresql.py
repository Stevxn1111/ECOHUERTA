import psycopg2
try: 
    connection = psycopg2.connect(
        host='localhost',
        user='postgres',
        password='123',
        database='ECOHUERTA'
    )
    
    print("Conexión exitosa a la base de datos PostgreSQL")
    cursor = connection.cursor()
    cursor.execute("SELECT version();")
    row = cursor.fetchone()
    print("Versión de la base de datos:", row)
    cursor.execute("SELECT * FROM plantas;")
    rows = cursor.fetchall()
    for row in rows:
        print(row)
except Exception as ex:
    print("Error al conectar a la base de datos:", ex)
finally:
    connection.close()
    print("Conexión cerrada")