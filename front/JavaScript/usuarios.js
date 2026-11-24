document.addEventListener("DOMContentLoaded", async () => {
  const sesionRaw = localStorage.getItem("ecohuerta_sesion");
  if (!sesionRaw) {
    window.location.href = "../Login/login.html";
    return;
  }

  let sesion = {};
  try {
    sesion = JSON.parse(sesionRaw) || {};
  } catch {
    localStorage.removeItem("ecohuerta_sesion");
    window.location.href = "../Login/login.html";
    return;
  }

  const rolSesion = String(sesion.rol || "").trim().toLowerCase();
  if (rolSesion !== "admin") {
    window.location.href = "ventanaPrincipal.html";
    return;
  }

  const tituloPaginaEl = document.getElementById("tituloPaginaUsuarios");
  const tituloPanelEl = document.getElementById("tituloPanelUsuarios");
  const subtituloPanelEl = document.getElementById("subtituloPanelUsuarios");
  const avatarEl = document.getElementById("avatarAdmin");
  const nombreEl = document.getElementById("nombreAdmin");
  const rolEl = document.getElementById("rolAdmin");
  const btnVolver = document.getElementById("btnVolverMapaAdmin");
  const filtroRol = document.getElementById("filtroRolUsuario");
  const btnNuevoUsuario = document.getElementById("btnNuevoUsuario");
  const contadorEl = document.getElementById("contadorUsuarios");
  const estadoEl = document.getElementById("estadoUsuarios");
  const tablaWrapper = document.getElementById("contenedorTablaUsuarios");
  const tbody = document.getElementById("cuerpoTablaUsuarios");

  const panelForm = document.getElementById("panelFormularioUsuario");
  const tituloForm = document.getElementById("tituloFormUsuario");
  const formUsuario = document.getElementById("formUsuarioAdmin");
  const inputId = document.getElementById("usuarioIdInput");
  const inputNombre = document.getElementById("usuarioNombreInput");
  const inputUsuario = document.getElementById("usuarioLoginInput");
  const inputClave = document.getElementById("usuarioClaveInput");
  const inputRol = document.getElementById("usuarioRolInput");
  const btnCancelarUsuario = document.getElementById("btnCancelarUsuario");
  const msgUsuario = document.getElementById("mensajeUsuarioAdmin");

  if (avatarEl && sesion.avatar) avatarEl.textContent = sesion.avatar;
  if (nombreEl) nombreEl.textContent = sesion.nombre || sesion.usuario || "Admin";
  if (rolEl) rolEl.textContent = rolSesion;

  btnVolver?.addEventListener("click", () => {
    window.location.href = "ventanaPrincipal.html";
  });

  const params = new URLSearchParams(window.location.search);
  const rolInicial = (params.get("rol") || "todos").trim().toLowerCase();
  if (filtroRol && ["todos", "estudiante", "docente", "admin"].includes(rolInicial)) {
    filtroRol.value = rolInicial;
  }

  function actualizarEncabezadoPorRol() {
    const actual = filtroRol ? filtroRol.value : "todos";
    if (actual === "estudiante") {
      if (tituloPaginaEl) tituloPaginaEl.textContent = "Estudiantes - ECOHUERTA";
      if (tituloPanelEl) {
        tituloPanelEl.innerHTML = `<span class="co">co</span> Gestión de Estudiantes`;
      }
      if (subtituloPanelEl) {
        subtituloPanelEl.textContent = "Consulta, crea, edita o elimina cuentas de estudiantes inscritos";
      }
      if (btnNuevoUsuario) btnNuevoUsuario.textContent = "+ Nuevo estudiante";
    } else if (actual === "docente") {
      if (tituloPaginaEl) tituloPaginaEl.textContent = "Docentes - ECOHUERTA";
      if (tituloPanelEl) {
        tituloPanelEl.innerHTML = `<span class="co">co</span> Gestión de Docentes`;
      }
      if (subtituloPanelEl) {
        subtituloPanelEl.textContent = "Consulta, crea, edita o elimina cuentas de docentes inscritos";
      }
      if (btnNuevoUsuario) btnNuevoUsuario.textContent = "+ Nuevo docente";
    } else {
      if (tituloPaginaEl) tituloPaginaEl.textContent = "Administración - ECOHUERTA";
      if (tituloPanelEl) {
        tituloPanelEl.innerHTML = `<span class="co">co</span> Panel de Administración`;
      }
      if (subtituloPanelEl) {
        subtituloPanelEl.textContent = "Gestión de usuarios y roles de ECOHUERTA";
      }
      if (btnNuevoUsuario) btnNuevoUsuario.textContent = "+ Nuevo usuario";
    }
  }

  function mostrarMensajeForm(texto, esError = false) {
    if (!msgUsuario) return;
    msgUsuario.textContent = texto;
    msgUsuario.className = esError ? "mensaje-admin error" : "mensaje-admin exito";
    msgUsuario.style.display = "block";
  }

  function ocultarMensajeForm() {
    if (!msgUsuario) return;
    msgUsuario.style.display = "none";
    msgUsuario.textContent = "";
  }

  function abrirFormularioCrear() {
    ocultarMensajeForm();
    const rolActivo =
      filtroRol && ["estudiante", "docente", "admin"].includes(filtroRol.value)
        ? filtroRol.value
        : "estudiante";

    if (tituloForm) {
      if (rolActivo === "estudiante") tituloForm.textContent = "Registrar nuevo estudiante";
      else if (rolActivo === "docente") tituloForm.textContent = "Registrar nuevo docente";
      else tituloForm.textContent = "Registrar nuevo usuario";
    }

    if (inputId) inputId.value = "";
    if (inputNombre) inputNombre.value = "";
    if (inputUsuario) inputUsuario.value = "";
    if (inputClave) {
      inputClave.value = "";
      inputClave.required = true;
      inputClave.placeholder = "Contraseña";
    }
    if (inputRol) inputRol.value = rolActivo;

    if (panelForm) {
      panelForm.style.display = "block";
      inputNombre?.focus();
    }
  }

  function abrirFormularioEditar(u) {
    ocultarMensajeForm();
    if (tituloForm) tituloForm.textContent = `Editar usuario #${u.id} (${u.usuario})`;
    if (inputId) inputId.value = String(u.id);
    if (inputNombre) inputNombre.value = u.nombre || "";
    if (inputUsuario) inputUsuario.value = u.usuario || "";
    if (inputClave) {
      inputClave.value = "";
      inputClave.required = false;
      inputClave.placeholder = "Nueva contraseña (opcional)";
    }
    if (inputRol) inputRol.value = u.rol || "estudiante";

    if (panelForm) {
      panelForm.style.display = "block";
      inputNombre?.focus();
    }
  }

  function cerrarFormulario() {
    ocultarMensajeForm();
    if (panelForm) panelForm.style.display = "none";
    formUsuario?.reset();
  }

  btnNuevoUsuario?.addEventListener("click", abrirFormularioCrear);
  btnCancelarUsuario?.addEventListener("click", cerrarFormulario);

  formUsuario?.addEventListener("submit", async (e) => {
    e.preventDefault();
    ocultarMensajeForm();

    const idEditando = inputId ? Number(inputId.value) : 0;
    const nombre = inputNombre?.value.trim() || "";
    const usuario = inputUsuario?.value.trim() || "";
    const clave = inputClave?.value || "";
    const rol = inputRol?.value || "estudiante";

    if (!nombre || !usuario || !rol) {
      mostrarMensajeForm("Nombre, usuario y rol son obligatorios.", true);
      return;
    }

    if (!idEditando && !clave.trim()) {
      mostrarMensajeForm("La contraseña es obligatoria al crear un usuario.", true);
      return;
    }

    const url = idEditando ? `/api/usuarios/${encodeURIComponent(idEditando)}` : "/api/usuarios";
    const method = idEditando ? "PUT" : "POST";
    const payload = idEditando
      ? { nombre, usuario, rol, ...(clave.trim() ? { clave } : {}) }
      : { nombre, usuario, clave, rol };

    try {
      const resp = await fetch(url, {
        method,
        headers: {
          "Content-Type": "application/json",
          "X-User-Role": "admin",
        },
        body: JSON.stringify(payload),
      });
      const data = await resp.json().catch(() => ({}));
      if (!resp.ok) {
        mostrarMensajeForm(data.error || "No se pudo guardar el usuario.", true);
        return;
      }

      cerrarFormulario();
      await cargarUsuarios();
    } catch {
      mostrarMensajeForm("Error de conexión al guardar el usuario.", true);
    }
  });

  async function cargarUsuarios() {
    actualizarEncabezadoPorRol();
    const rolSeleccionado = filtroRol ? filtroRol.value : "todos";
    if (estadoEl) {
      estadoEl.style.display = "block";
      estadoEl.textContent = "Cargando usuarios...";
    }
    if (tablaWrapper) tablaWrapper.style.display = "none";

    try {
      const resp = await fetch(`/api/usuarios?rol=${encodeURIComponent(rolSeleccionado)}`, {
        cache: "no-store",
        headers: {
          "X-User-Role": "admin",
        },
      });
      const data = await resp.json().catch(() => []);
      if (!resp.ok) {
        throw new Error(data.error || "Error al consultar usuarios");
      }

      const lista = Array.isArray(data) ? data : [];
      if (contadorEl) {
        contadorEl.textContent = `${lista.length} ${lista.length === 1 ? "usuario" : "usuarios"}`;
      }

      if (!lista.length) {
        if (estadoEl) {
          estadoEl.style.display = "block";
          estadoEl.textContent = "No hay usuarios registrados con este filtro.";
        }
        return;
      }

      if (tbody) {
        tbody.innerHTML = "";
        lista.forEach((u) => {
          const tr = document.createElement("tr");
          const esMiUsuario =
            (sesion.id && Number(sesion.id) === Number(u.id)) ||
            (sesion.usuario && sesion.usuario === u.usuario);

          tr.innerHTML = `
            <td>${escaparHTML(u.id)}</td>
            <td>${escaparHTML(u.nombre)}</td>
            <td>${escaparHTML(u.usuario)}</td>
            <td>
              <select class="select-rol-usuario" data-id="${escaparHTML(u.id)}">
                <option value="estudiante" ${u.rol === "estudiante" ? "selected" : ""}>estudiante</option>
                <option value="docente" ${u.rol === "docente" ? "selected" : ""}>docente</option>
                <option value="admin" ${u.rol === "admin" ? "selected" : ""}>admin</option>
              </select>
            </td>
            <td>
              <button type="button" class="btn-editar-planta btn-editar-usuario">Editar</button>
              <button type="button" class="btn-eliminar-planta btn-borrar-usuario" ${
                esMiUsuario ? "disabled title='No puedes eliminar tu propia sesión activa'" : ""
              }>
                Eliminar
              </button>
            </td>
          `;

          tr.querySelector(".select-rol-usuario")?.addEventListener("change", async (e) => {
            const nuevoRol = e.target.value;
            try {
              const rPut = await fetch(`/api/usuarios/${encodeURIComponent(u.id)}`, {
                method: "PUT",
                headers: {
                  "Content-Type": "application/json",
                  "X-User-Role": "admin",
                },
                body: JSON.stringify({
                  nombre: u.nombre,
                  usuario: u.usuario,
                  rol: nuevoRol,
                }),
              });
              const dPut = await rPut.json().catch(() => ({}));
              if (!rPut.ok) {
                alert(dPut.error || "No se pudo actualizar el rol.");
                await cargarUsuarios();
                return;
              }
              await cargarUsuarios();
            } catch {
              alert("Error de conexión al cambiar rol.");
            }
          });

          tr.querySelector(".btn-editar-usuario")?.addEventListener("click", () => {
            abrirFormularioEditar(u);
          });

          tr.querySelector(".btn-borrar-usuario")?.addEventListener("click", async () => {
            if (esMiUsuario) return;
            const ok = window.confirm(`¿Eliminar al usuario "${u.usuario}"?`);
            if (!ok) return;
            try {
              const rDel = await fetch(`/api/usuarios/${encodeURIComponent(u.id)}`, {
                method: "DELETE",
                headers: {
                  "X-User-Role": "admin",
                },
              });
              const dDel = await rDel.json().catch(() => ({}));
              if (!rDel.ok) {
                alert(dDel.error || "No se pudo eliminar el usuario.");
                return;
              }
              await cargarUsuarios();
            } catch {
              alert("Error de conexión al eliminar usuario.");
            }
          });

          tbody.appendChild(tr);
        });
      }

      if (estadoEl) estadoEl.style.display = "none";
      if (tablaWrapper) tablaWrapper.style.display = "block";
    } catch (err) {
      console.error("Error cargando usuarios:", err);
      if (estadoEl) {
        estadoEl.style.display = "block";
        estadoEl.textContent = `No se pudieron cargar los usuarios (${err.message}).`;
      }
    }
  }

  filtroRol?.addEventListener("change", cargarUsuarios);

  await cargarUsuarios();
});

function escaparHTML(texto) {
  return String(texto ?? "")
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#39;");
}