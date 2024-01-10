$('#seguir').on('click', seguir);
$('#parar-de-seguir').on('click', pararDeSeguir);

function seguir(event) {
  event.preventDefault();
  const usuarioId = $(this).data('usuario-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/usuarios/${usuarioId}/seguir`,
    method: "POST",
  }).done(function() {
    window.location = `/usuarios/${usuarioId}`
  }).fail(function() {
    Swal.fire('Erro!', 'Falha ao seguir usuário!', 'error')
    $(this).prop('disabled', false);
  })
}

function pararDeSeguir(event) {
  event.preventDefault();
  const usuarioId = $(this).data('usuario-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/usuarios/${usuarioId}/parar-de-seguir`,
    method: "POST",
  }).done(function() {
    window.location = `/usuarios/${usuarioId}`
  }).fail(function() {
    Swal.fire('Erro!', 'Falha ao parar de seguir usuário!', 'error')
    $(this).prop('disabled', false);
  })
}