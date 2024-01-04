$('#nova-publicacao').on('submit', criarPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);

function criarPublicacao(event) {
  event.preventDefault();

  $.ajax({
    url: "/publicacoes",
    method: "POST",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val(),
    }
  }).done(function() {
    window.location = "/home";
  }).fail(function() {
    alert("Falha ao criar publicação");
  })
}

function curtirPublicacao(event) {
  event.preventDefault();

  const elementoCriado = $(event.target);
  const publicacaoId = elementoCriado.closest('div').data('publicacao-id');

  elementoCriado.prop('disabled', true);
  $.ajax({
    url: `/publicacoes/${publicacaoId}/curtir`,
    method: "POST",
  }).done(function() {
    const contadorDeCurtidas = elementoCriado.next('span');
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

    elementoCriado.addClass('descurtir-publicacao');
    elementoCriado.addClass('text-danger');
    elementoCriado.removeClass('curtir-publicacao');

    contadorDeCurtidas.text(quantidadeDeCurtidas + 1);
  }).fail(function() {
    alert("Falha ao curtir publicação");
  }).always(function() {
    elementoCriado.prop('disabled', false)
  });
}

function descurtirPublicacao(event) {
  event.preventDefault();

  const elementoCriado = $(event.target);
  const publicacaoId = elementoCriado.closest('div').data('publicacao-id');

  elementoCriado.prop('disabled', true);
  $.ajax({
    url: `/publicacoes/${publicacaoId}/descurtir`,
    method: "POST",
  }).done(function() {
    const contadorDeCurtidas = elementoCriado.next('span');
    const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

    elementoCriado.removeClass('descurtir-publicacao');
    elementoCriado.removeClass('text-danger');
    elementoCriado.addClass('curtir-publicacao');

    contadorDeCurtidas.text(quantidadeDeCurtidas - 1);
  }).fail(function() {
    alert("Falha ao descurtir publicação");
  }).always(function() {
    elementoCriado.prop('disabled', false)
  });
}