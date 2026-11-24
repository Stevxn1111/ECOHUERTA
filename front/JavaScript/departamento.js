async function iniciarDepartamento() {
  if (window.__ecohuertaDeptoInicializado) return;
  window.__ecohuertaDeptoInicializado = true;

  const estadoEl = document.getElementById("estadoCarga");
  const gridEl = document.getElementById("gridPlantas");

  const mostrarErrorVisible = (mensaje) => {
    if (estadoEl) {
      estadoEl.style.display = "block";
      estadoEl.textContent = mensaje;
    }
  };

  try {
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

    const rolDetectado = String(
      sesion.rol || sesion.role || sesion.tipo_usuario || "estudiante"
    )
      .trim()
      .toLowerCase();
    const puedeGestionar = rolDetectado === "docente" || rolDetectado === "admin";

    const avatarEl = document.getElementById("avatarDepto");
    const nombreUsuarioEl = document.getElementById("nombreDeptoUsuario");
    const rolUsuarioEl = document.getElementById("rolDeptoUsuario");
    const btnVolverMapa = document.getElementById("btnVolverMapa");
    const tituloDeptoEl = document.getElementById("tituloDepartamento");
    const subtituloDeptoEl = document.getElementById("subtituloDepartamento");
    const filtroTipoEl = document.getElementById("filtroTipo");
    const contadorEl = document.getElementById("contadorPlantas");
    const btnNuevaPlanta = document.getElementById("btnNuevaPlanta");

    const modalEl = document.getElementById("modalPlanta");
    const tituloModalEl = document.getElementById("tituloModalPlanta");
    const formPlanta = document.getElementById("formPlanta");
    const inputId = document.getElementById("plantaIdInput");
    const inputNombre = document.getElementById("plantaNombreInput");
    const inputTipo = document.getElementById("plantaTipoInput");
    const inputDepto = document.getElementById("plantaDeptoInput");
    const inputImagen = document.getElementById("plantaImagenInput");
    const inputDesc = document.getElementById("plantaDescInput");
    const inputCarac = document.getElementById("plantaCaracInput");
    const inputCuidados = document.getElementById("plantaCuidadosInput");
    const inputBeneficios = document.getElementById("plantaBeneficiosInput");
    const errorModalEl = document.getElementById("errorModalPlanta");
    const btnCerrarModal = document.getElementById("btnCerrarModal");
    const btnCancelarModal = document.getElementById("btnCancelarModal");

    if (avatarEl && sesion.avatar) avatarEl.textContent = sesion.avatar;
    if (nombreUsuarioEl) nombreUsuarioEl.textContent = sesion.nombre || sesion.usuario || "Usuario";
    if (rolUsuarioEl) rolUsuarioEl.textContent = rolDetectado;

    if (btnNuevaPlanta) {
      btnNuevaPlanta.style.setProperty(
        "display",
        puedeGestionar ? "inline-block" : "none",
        "important"
      );
    }

    if (btnVolverMapa && !btnVolverMapa.dataset.bound) {
      btnVolverMapa.dataset.bound = "1";
      btnVolverMapa.addEventListener("click", () => {
        window.location.href = "ventanaPrincipal.html";
      });
    }

    const params = new URLSearchParams(window.location.search);
    let slug = (params.get("slug") || "").trim().toLowerCase();

    let deptoGuardado = null;
    try {
      deptoGuardado = JSON.parse(localStorage.getItem("ecohuerta_departamento_actual") || "null");
    } catch {
      deptoGuardado = null;
    }

    if (!slug && deptoGuardado && deptoGuardado.slug) {
      slug = String(deptoGuardado.slug).toLowerCase();
    }

    if (!slug) {
      window.location.href = "ventanaPrincipal.html";
      return;
    }

    let nombreDepartamento =
      deptoGuardado && deptoGuardado.slug === slug && deptoGuardado.nombre
        ? deptoGuardado.nombre
        : slug.replaceAll("_", " ").replace(/\b\w/g, (c) => c.toUpperCase());

    let idDepartamentoActual =
      deptoGuardado && deptoGuardado.slug === slug ? deptoGuardado.id_departamento : 0;

    if (tituloDeptoEl) tituloDeptoEl.textContent = nombreDepartamento;
    if (subtituloDeptoEl) {
      subtituloDeptoEl.textContent = `Especies de flora registradas en ${nombreDepartamento}`;
    }

    const [respDeptos, respTipos] = await Promise.all([
      fetch("/api/departamentos", { cache: "no-store" }).catch(() => null),
      fetch("/api/tipos-planta", { cache: "no-store" }).catch(() => null),
    ]);

    let listaDepartamentos = [];
    if (respDeptos && respDeptos.ok) {
      const dataDeptos = await respDeptos.json().catch(() => []);
      if (Array.isArray(dataDeptos)) {
        listaDepartamentos = dataDeptos;
        const encontrado = listaDepartamentos.find((d) => d.slug === slug);
        if (encontrado) {
          nombreDepartamento = encontrado.nombre;
          idDepartamentoActual = encontrado.id_departamento;
          localStorage.setItem("ecohuerta_departamento_actual", JSON.stringify(encontrado));
          if (tituloDeptoEl) tituloDeptoEl.textContent = nombreDepartamento;
          if (subtituloDeptoEl) {
            subtituloDeptoEl.textContent = `Especies de flora registradas en ${nombreDepartamento}`;
          }
        }
      }
    }

    let listaTipos = [];
    if (respTipos && respTipos.ok) {
      const dataTipos = await respTipos.json().catch(() => []);
      if (Array.isArray(dataTipos)) {
        listaTipos = dataTipos;
      }
    }

    if (filtroTipoEl) {
      filtroTipoEl.innerHTML = '<option value="">Todos los tipos</option>';
      listaTipos.forEach((t) => {
        const option = document.createElement("option");
        option.value = t.id_tipo;
        option.textContent = t.nombre;
        filtroTipoEl.appendChild(option);
      });
    }

    function obtenerIdPlanta(planta) {
      return planta?.id_planta ?? planta?.id ?? "";
    }

    function poblarSelectsModal() {
      if (inputTipo) {
        inputTipo.innerHTML = "";
        listaTipos.forEach((t) => {
          const opt = document.createElement("option");
          opt.value = t.id_tipo;
          opt.textContent = t.nombre;
          inputTipo.appendChild(opt);
        });
      }
      if (inputDepto) {
        inputDepto.innerHTML = "";
        listaDepartamentos.forEach((d) => {
          const opt = document.createElement("option");
          opt.value = d.id_departamento;
          opt.textContent = d.nombre;
          inputDepto.appendChild(opt);
        });
      }
    }

    function mostrarModal() {
      if (modalEl) {
        modalEl.style.setProperty("display", "flex", "important");
      }
    }

    function cerrarModal() {
      if (modalEl) {
        modalEl.style.setProperty("display", "none", "important");
      }
    }

    function abrirModalCrear() {
      poblarSelectsModal();
      if (tituloModalEl) tituloModalEl.textContent = "Agregar nueva planta";
      if (inputId) inputId.value = "";
      if (inputNombre) inputNombre.value = "";
      if (inputImagen) inputImagen.value = "";
      if (inputDesc) inputDesc.value = "";
      if (inputCarac) inputCarac.value = "";
      if (inputCuidados) inputCuidados.value = "";
      if (inputBeneficios) inputBeneficios.value = "";
      if (inputDepto && idDepartamentoActual) {
        inputDepto.value = String(idDepartamentoActual);
      }
      if (errorModalEl) errorModalEl.style.display = "none";
      mostrarModal();
    }

    function abrirModalEditar(planta) {
      poblarSelectsModal();
      if (tituloModalEl) tituloModalEl.textContent = `Editar: ${planta.nombre || "Planta"}`;
      if (inputId) inputId.value = obtenerIdPlanta(planta);
      if (inputNombre) inputNombre.value = planta.nombre || "";
      if (inputTipo) inputTipo.value = String(planta.id_tipo || "");
      if (inputDepto) {
        inputDepto.value = String(planta.id_departamento || idDepartamentoActual || "");
      }
      if (inputImagen) inputImagen.value = planta.imagen || "";
      if (inputDesc) inputDesc.value = planta.descripcion || "";
      if (inputCarac) inputCarac.value = planta.caracteristicas || "";
      if (inputCuidados) inputCuidados.value = planta.cuidados || "";
      if (inputBeneficios) inputBeneficios.value = planta.beneficios || "";
      if (errorModalEl) errorModalEl.style.display = "none";
      mostrarModal();
    }

    function irAFichaPlanta(planta) {
      const idPlanta = obtenerIdPlanta(planta);
      localStorage.setItem("ecohuerta_planta_actual", JSON.stringify(planta));
      const destino = idPlanta
        ? `planta.html?id=${encodeURIComponent(idPlanta)}`
        : "planta.html";
      window.location.href = destino;
    }

    async function eliminarPlanta(planta) {
      const idPlanta = obtenerIdPlanta(planta);
      if (!idPlanta) {
        alert("No se encontró el ID de la planta.");
        return;
      }

      const confirmado = window.confirm(
        `¿Seguro que deseas eliminar "${planta.nombre || "esta planta"}" de ${nombreDepartamento}?`
      );
      if (!confirmado) return;

      try {
        const resp = await fetch(`/api/plantas/${encodeURIComponent(idPlanta)}`, {
          method: "DELETE",
          headers: {
            "X-User-Role": rolDetectado,
          },
        });
        const data = await resp.json().catch(() => ({}));
        if (!resp.ok) {
          alert(data.error || "No se pudo eliminar la planta.");
          return;
        }
        await cargarPlantas(filtroTipoEl ? filtroTipoEl.value : "");
      } catch (err) {
        console.error("Error al eliminar planta:", err);
        alert("Error de conexión al intentar eliminar la planta.");
      }
    }

    function renderizarPlantas(plantas) {
      const total = Array.isArray(plantas) ? plantas.length : 0;
      if (contadorEl) {
        contadorEl.textContent = `${total} ${total === 1 ? "planta" : "plantas"}`;
      }

      if (!gridEl) return;
      gridEl.innerHTML = "";

      if (!total) {
        mostrarErrorVisible(
          `Aún no hay plantas registradas para ${nombreDepartamento} con este filtro.`
        );
        return;
      }

      if (estadoEl) estadoEl.style.display = "none";

      plantas.forEach((planta) => {
        const card = document.createElement("article");
        card.className = "tarjeta-planta";

        const nombrePlanta = planta?.nombre || "Planta sin nombre";
        const descPlanta = planta?.descripcion || "Sin descripción disponible.";
        const tipoNombre = planta?.tipo_nombre || "";
        const imagenUrl = planta?.imagen || "";

        const badgeTipo = tipoNombre
          ? `<span class="badge-tipo">${escaparHTML(tipoNombre)}</span>`
          : "";

        const imgHTML = imagenUrl
          ? `<img src="${escaparHTML(imagenUrl)}" alt="${escaparHTML(nombrePlanta)}" class="img-planta">`
          : `<div class="img-planta-placeholder">🌿</div>`;

        const accionesDocenteHTML = puedeGestionar
          ? `
            <div class="acciones-docente-tarjeta">
              <button type="button" class="btn-editar-planta">Editar</button>
              <button type="button" class="btn-eliminar-planta">Eliminar</button>
            </div>
          `
          : "";

        card.innerHTML = `
          ${imgHTML}
          <div class="contenido-planta">
            <div class="cabecera-tarjeta">
              <h3>${escaparHTML(nombrePlanta)}</h3>
              ${badgeTipo}
            </div>
            <p class="desc-planta">${escaparHTML(descPlanta)}</p>
            <button type="button" class="btn-ver-planta">Ver ficha completa</button>
            ${accionesDocenteHTML}
          </div>
        `;

        card.querySelector(".btn-ver-planta")?.addEventListener("click", (e) => {
          e.preventDefault();
          irAFichaPlanta(planta);
        });

        if (puedeGestionar) {
          card.querySelector(".btn-editar-planta")?.addEventListener("click", (e) => {
            e.preventDefault();
            abrirModalEditar(planta);
          });
          card.querySelector(".btn-eliminar-planta")?.addEventListener("click", (e) => {
            e.preventDefault();
            eliminarPlanta(planta);
          });
        }

        gridEl.appendChild(card);
      });
    }

    async function cargarPlantas(tipoId = "") {
      if (estadoEl) {
        estadoEl.style.display = "block";
        estadoEl.textContent = "Cargando plantas...";
      }
      if (gridEl) gridEl.innerHTML = "";

      const query = tipoId ? `?tipo=${encodeURIComponent(tipoId)}` : "";
      const resp = await fetch(
        `/api/departamentos/${encodeURIComponent(slug)}/plantas${query}`,
        { cache: "no-store" }
      );
      const data = await resp.json().catch(() => null);

      if (!resp.ok) {
        throw new Error((data && data.error) || `HTTP ${resp.status}`);
      }

      renderizarPlantas(Array.isArray(data) ? data : []);
    }

    if (btnNuevaPlanta && !btnNuevaPlanta.dataset.bound) {
      btnNuevaPlanta.dataset.bound = "1";
      btnNuevaPlanta.addEventListener("click", abrirModalCrear);
    }

    if (btnCerrarModal && !btnCerrarModal.dataset.bound) {
      btnCerrarModal.dataset.bound = "1";
      btnCerrarModal.addEventListener("click", cerrarModal);
    }

    if (btnCancelarModal && !btnCancelarModal.dataset.bound) {
      btnCancelarModal.dataset.bound = "1";
      btnCancelarModal.addEventListener("click", cerrarModal);
    }

    if (formPlanta && !formPlanta.dataset.bound) {
      formPlanta.dataset.bound = "1";
      formPlanta.addEventListener("submit", async (e) => {
        e.preventDefault();
        if (errorModalEl) errorModalEl.style.display = "none";

        const idEditando = inputId ? Number(inputId.value) : 0;
        const payload = {
          nombre: inputNombre?.value.trim() || "",
          id_tipo: Number(inputTipo?.value || 0),
          id_departamento: Number(inputDepto?.value || 0),
          imagen: inputImagen?.value.trim() || "",
          descripcion: inputDesc?.value.trim() || "",
          caracteristicas: inputCarac?.value.trim() || "",
          cuidados: inputCuidados?.value.trim() || "",
          beneficios: inputBeneficios?.value.trim() || "",
        };

        const url = idEditando ? `/api/plantas/${idEditando}` : "/api/plantas";
        const method = idEditando ? "PUT" : "POST";

        try {
          const resp = await fetch(url, {
            method,
            headers: {
              "Content-Type": "application/json",
              "X-User-Role": rolDetectado,
            },
            body: JSON.stringify(payload),
          });
          const data = await resp.json().catch(() => ({}));
          if (!resp.ok) {
            if (errorModalEl) {
              errorModalEl.textContent = data.error || "No se pudo guardar la planta.";
              errorModalEl.style.display = "block";
            }
            return;
          }

          cerrarModal();
          await cargarPlantas(filtroTipoEl ? filtroTipoEl.value : "");
        } catch (err) {
          console.error("Error guardando planta:", err);
          if (errorModalEl) {
            errorModalEl.textContent = "Error de red al guardar la planta.";
            errorModalEl.style.display = "block";
          }
        }
      });
    }

    if (filtroTipoEl && !filtroTipoEl.dataset.bound) {
      filtroTipoEl.dataset.bound = "1";
      filtroTipoEl.addEventListener("change", () => {
        cargarPlantas(filtroTipoEl.value).catch((err) => {
          console.error("Error al filtrar plantas:", err);
          mostrarErrorVisible(`No se pudieron filtrar las plantas (${err.message}).`);
        });
      });
    }

    await cargarPlantas();
  } catch (err) {
    console.error("Error fatal en departamento.js:", err);
    mostrarErrorVisible(`Ocurrió un error al cargar el departamento: ${err.message}`);
  }
}

function escaparHTML(texto) {
  return String(texto ?? "")
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#39;");
}

if (document.readyState === "loading") {
  document.addEventListener("DOMContentLoaded", () => {
    iniciarDepartamento();
  });
} else {
  iniciarDepartamento();
}

window.addEventListener("pageshow", () => {
  window.__ecohuertaDeptoInicializado = false;
  iniciarDepartamento();
});