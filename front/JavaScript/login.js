let avatarSeleccionado = "🦊";

document.querySelectorAll(".avatar-option").forEach((btn) => {
  btn.addEventListener("click", () => {
    document.querySelectorAll(".avatar-option").forEach((b) => {
      b.classList.remove("selected");
      b.setAttribute("aria-pressed", "false");
    });
    btn.classList.add("selected");
    btn.setAttribute("aria-pressed", "true");
    avatarSeleccionado = btn.textContent.trim();
  });
});

const loginForm = document.getElementById("loginForm");
const mensajeLogin = document.getElementById("mensajeLogin");

loginForm.addEventListener("submit", async (e) => {
  e.preventDefault();

  const usuario = document.getElementById("user").value.trim();
  const clave = document.getElementById("pass").value;

  if (!usuario || !clave) {
    mensajeLogin.textContent = "Completa tu usuario y contraseña.";
    return;
  }

  mensajeLogin.textContent = "Entrando...";

  try {
    const resp = await fetch("/api/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ usuario, clave })
    });

    const data = await resp.json();

    if (!resp.ok) {
      mensajeLogin.textContent = data.error || "Usuario o contraseña incorrectos.";
      return;
    }

    localStorage.setItem(
      "ecohuerta_sesion",
      JSON.stringify({
        ...data.usuario,
        avatar: avatarSeleccionado
      })
    );

    window.location.href = "../templates/ventanaPrincipal.html";
  } catch (err) {
    mensajeLogin.textContent = "Error de conexión con el servidor.";
  }
});