$('#nova-publicacao').on('submit', criarPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);
$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('.deletar-publicacao').on('click', deletarPublicacao);

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
    Swal.fire('Erro!', 'Falha ao criar publicação', 'error');
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
    Swal.fire('Erro!', 'Falha ao curtir publicação', 'error');
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
    Swal.fire('Erro!', 'Falha ao descurtir publicação', 'error');
  }).always(function() {
    elementoCriado.prop('disabled', false)
  });
}

function atualizarPublicacao(event) {
  event.preventDefault();
  $(this).prop('disabled', true);
  const publicacaoId = $(this).data('publicacao-id')

  $.ajax({
    url: `/publicacoes/${publicacaoId}`,
    method: "PUT",
    data: {
      titulo: $('#titulo').val(),
      conteudo: $('#conteudo').val()
    }
  }).done(function() {
    Swal.fire('Sucesso!', 'Publicação alterada!', 'success').then(function() {
      window.location = "/home";
    });
  }).fail(function() {
    Swal.fire('Erro!', 'Falha ao atualizar publicação!', 'error');
  }).always(function() {
    $('#atualizar-publicacao').prop('disabled', false);
  })
}

function deletarPublicacao(event) {
  event.preventDefault();
  Swal.fire({
    title: "Atenção!",
    text: "Tem certeza que deseja excluir a publicação?",
    showCancelButton: true,
    cancelButtonText: "Cancelar",
    icon: "warning"
  }).then(function(confirmacao) {
    if (!confirmacao.value) return;
    const elementoCriado = $(event.target);
    const publicacao = elementoCriado.closest('div');
    const publicacaoId = elementoCriado.closest('div').data('publicacao-id');

    elementoCriado.prop('disabled', true);
    $.ajax({
      url: `publicacoes/${publicacaoId}`,
      method: "DELETE"
    }).done(function() {
      publicacao.fadeOut("slow", function() {
        $(this).remove();
      })
    }).fail(function() {
      Swal.fire('Erro!', 'Falha ao excluir publicação!', 'error');
    })
  })
}
