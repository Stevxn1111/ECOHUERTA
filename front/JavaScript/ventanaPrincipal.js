document.addEventListener("DOMContentLoaded", async () => {
  const sesionRaw = localStorage.getItem("ecohuerta_sesion");
  if (!sesionRaw) {
    window.location.href = "../Login/login.html";
    return;
  }

  let sesion;
  try {
    sesion = JSON.parse(sesionRaw);
  } catch {
    localStorage.removeItem("ecohuerta_sesion");
    window.location.href = "../Login/login.html";
    return;
  }

  const nombreEl = document.getElementById("nombreUsuario");
  const rolEl = document.getElementById("rolUsuario");
  const avatarEl = document.getElementById("avatarUsuario");
  const saludoRolEl = document.getElementById("saludoRol");
  const accionesAdminEl = document.getElementById("accionesAdmin");
  const btnEstudiantes = document.getElementById("btnEstudiantes");
  const btnDocentes = document.getElementById("btnDocentes");
  const btnCerrarSesion = document.getElementById("btnCerrarSesion");

  if (nombreEl) nombreEl.textContent = sesion.nombre || sesion.usuario || "Usuario";
  if (rolEl) rolEl.textContent = sesion.rol || "estudiante";
  if (avatarEl && sesion.avatar) avatarEl.textContent = sesion.avatar;
  if (saludoRolEl) {
    saludoRolEl.textContent = `HOLA ${(sesion.rol || "explorador").toUpperCase()}`;
  }

  if (sesion.rol === "admin" && accionesAdminEl) {
    accionesAdminEl.style.display = "flex";
    btnEstudiantes?.addEventListener("click", () => {
      window.location.href = "usuarios.html?rol=estudiante";
    });
    btnDocentes?.addEventListener("click", () => {
      window.location.href = "usuarios.html?rol=docente";
    });
  }

  btnCerrarSesion?.addEventListener("click", () => {
    localStorage.removeItem("ecohuerta_sesion");
    window.location.href = "../Login/login.html";
  });

  // Crear tooltip flotante instantáneo para el mapa
  const tooltip = document.createElement("div");
  tooltip.className = "tooltip-mapa";
  document.body.appendChild(tooltip);

  let departamentosPorSlug = {};
  try {
    const resp = await fetch("/api/departamentos");
    if (resp.ok) {
      const lista = await resp.json();
      lista.forEach((d) => {
        departamentosPorSlug[d.slug] = d;
      });
    }
  } catch (err) {
    console.error("Error cargando departamentos:", err);
  }

  document.querySelectorAll(".zona").forEach((btn) => {
    const slug = btn.dataset.slug;
    const depto = departamentosPorSlug[slug];
    const nombreDepto = (depto && depto.nombre) ? depto.nombre : (btn.title || slug);

    // Guardamos en data-nombre y removemos title nativo para evitar doble tooltip
    btn.dataset.nombre = nombreDepto;
    btn.removeAttribute("title");

    btn.addEventListener("mouseenter", (e) => {
      tooltip.textContent = btn.dataset.nombre;
      tooltip.style.left = `${e.clientX}px`;
      tooltip.style.top = `${e.clientY}px`;
      tooltip.classList.add("visible");
    });

    btn.addEventListener("mousemove", (e) => {
      tooltip.style.left = `${e.clientX}px`;
      tooltip.style.top = `${e.clientY}px`;
    });

    btn.addEventListener("mouseleave", () => {
      tooltip.classList.remove("visible");
    });

    btn.addEventListener("click", () => {
      const departamentoSeleccionado = depto || {
        slug,
        nombre: nombreDepto
      };
      localStorage.setItem(
        "ecohuerta_departamento_actual",
        JSON.stringify(departamentoSeleccionado)
      );
      window.location.href = `departamento.html?slug=${encodeURIComponent(slug)}`;
    });
  });
});