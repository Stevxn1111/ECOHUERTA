const form = document.getElementById("registroForm");
const mensaje = document.getElementById("mensaje");

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  mensaje.textContent = "Registrando...";

  const datos = {
    nombre: document.getElementById("nombre").value.trim(),
    usuario: document.getElementById("usuario").value.trim(),
    clave: document.getElementById("clave").value,
    rol: document.getElementById("rol").value
  };

  try {
    const resp = await fetch("/api/registro", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(datos)
    });

    const data = await resp.json();

    if (!resp.ok) {
      mensaje.textContent = data.error || "No se pudo completar el registro.";
      return;
    }

    mensaje.textContent = "¡Registro exitoso! Redirigiendo al login...";
    setTimeout(() => {
      window.location.href = "login.html";
    }, 1000);
  } catch (err) {
    mensaje.textContent = "Error de conexión con el servidor.";
  }
});