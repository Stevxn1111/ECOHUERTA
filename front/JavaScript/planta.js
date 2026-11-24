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

  const avatarEl = document.getElementById("avatarPlanta");
  const nombreUsuarioEl = document.getElementById("nombreUsuarioPlanta");
  const rolUsuarioEl = document.getElementById("rolUsuarioPlanta");
  const btnVolverDepto = document.getElementById("btnVolverDepto");

  if (avatarEl && sesion.avatar) avatarEl.textContent = sesion.avatar;
  if (nombreUsuarioEl) nombreUsuarioEl.textContent = sesion.nombre || sesion.usuario || "Usuario";
  if (rolUsuarioEl) rolUsuarioEl.textContent = sesion.rol || "estudiante";

  const params = new URLSearchParams(window.location.search);
  const idParam = params.get("id");

  let plantaGuardada = null;
  try {
    plantaGuardada = JSON.parse(localStorage.getItem("ecohuerta_planta_actual") || "null");
  } catch {
    plantaGuardada = null;
  }

  let deptoGuardado = null;
  try {
    deptoGuardado = JSON.parse(localStorage.getItem("ecohuerta_departamento_actual") || "null");
  } catch {
    deptoGuardado = null;
  }

  btnVolverDepto?.addEventListener("click", () => {
    const slugDestino =
      (deptoGuardado && deptoGuardado.slug) ||
      (plantaGuardada && plantaGuardada.departamento_slug) ||
      "";
    if (slugDestino) {
      window.location.href = `departamento.html?slug=${encodeURIComponent(slugDestino)}`;
    } else if (window.history.length > 1) {
      window.history.back();
    } else {
      window.location.href = "ventanaPrincipal.html";
    }
  });

  let planta = null;
  if (idParam) {
    try {
      const resp = await fetch(`/api/plantas/${encodeURIComponent(idParam)}`, {
        cache: "no-store",
      });
      if (resp.ok) {
        planta = await resp.json();
      }
    } catch (err) {
      console.error("Error consultando planta por ID:", err);
    }
  }

  if (!planta && plantaGuardada) {
    planta = plantaGuardada;
  }

  const estadoEl = document.getElementById("estadoPlanta");
  const detalleEl = document.getElementById("detallePlanta");

  if (!planta) {
    if (estadoEl) {
      estadoEl.textContent = "No se encontró la información de la planta seleccionada.";
    }
    return;
  }

  const nombre = planta.nombre || "Planta sin nombre";
  const tipo = planta.tipo_nombre || "General";
  const departamento = planta.departamento_nombre || (deptoGuardado && deptoGuardado.nombre) || "Colombia";

  document.title = `${nombre} - ECOHUERTA`;
  const subtituloEl = document.getElementById("subtituloPlanta");
  if (subtituloEl) {
    subtituloEl.textContent = `${tipo} · ${departamento}`;
  }

  const nombreEl = document.getElementById("nombrePlantaDetalle");
  const tipoEl = document.getElementById("tipoPlantaDetalle");
  const deptoEl = document.getElementById("deptoPlantaDetalle");
  const descEl = document.getElementById("descPlantaDetalle");
  const caracEl = document.getElementById("caracPlantaDetalle");
  const cuidadosEl = document.getElementById("cuidadosPlantaDetalle");
  const beneficiosEl = document.getElementById("beneficiosPlantaDetalle");
  const imgEl = document.getElementById("imgPlantaDetalle");
  const placeholderEl = document.getElementById("placeholderPlantaDetalle");

  if (nombreEl) nombreEl.textContent = nombre;
  if (tipoEl) tipoEl.textContent = tipo;
  if (deptoEl) deptoEl.textContent = departamento;
  if (descEl) descEl.textContent = planta.descripcion || "Sin descripción registrada.";
  if (caracEl) caracEl.textContent = planta.caracteristicas || "Sin características registradas.";
  if (cuidadosEl) cuidadosEl.textContent = planta.cuidados || "Sin cuidados registrados.";
  if (beneficiosEl) beneficiosEl.textContent = planta.beneficios || "Sin beneficios registrados.";

  if (planta.imagen && imgEl) {
    const rutaImg = planta.imagen.startsWith("http") || planta.imagen.startsWith("/")
      ? planta.imagen
      : `../img/${planta.imagen}`;
    imgEl.src = rutaImg;
    imgEl.alt = nombre;
    imgEl.style.display = "block";
    if (placeholderEl) placeholderEl.style.display = "none";

    imgEl.onerror = () => {
      imgEl.style.display = "none";
      if (placeholderEl) placeholderEl.style.display = "flex";
    };
  }

  // Recursos educativos sugeridos según el nombre o tipo de planta
  const seccionRecursos = document.getElementById("seccionRecursos");
  const listaRecursos = document.getElementById("listaRecursos");
  if (seccionRecursos && listaRecursos) {
    const consulta = encodeURIComponent(`${nombre} ${departamento} flora colombia`);
    const recursos = [
      {
        titulo: `Buscar fichas y registros de ${nombre} en GBIF`,
        url: `https://www.gbif.org/es/species/search?q=${encodeURIComponent(nombre)}`,
      },
      {
        titulo: `Explorar información botánica de ${nombre}`,
        url: `https://es.wikipedia.org/wiki/Special:Search?search=${consulta}`,
      },
    ];

    listaRecursos.innerHTML = "";
    recursos.forEach((r) => {
      const link = document.createElement("a");
      link.href = r.url;
      link.target = "_blank";
      link.rel = "noopener noreferrer";
      link.className = "recurso-link";
      link.textContent = r.titulo;
      listaRecursos.appendChild(link);
    });
    seccionRecursos.style.display = "block";
  }

  if (estadoEl) estadoEl.style.display = "none";
  if (detalleEl) detalleEl.style.display = "grid";
});