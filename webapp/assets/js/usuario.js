$('#seguir').on('click', seguir);
$('#parar-de-seguir').on('click', pararDeSeguir);
$('#editar-usuario').on('submit', editar);
$('#atualizar-senha').on('submit', atualizarSenha);
$('#deletar-usuario').on('click', deletarUsuario);

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

function editar(event) {
  event.preventDefault();

  $.ajax({
    url: "/editar-usuario",
    method: "PUT",
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
    }
  }).done(function() {
    Swal.fire('Sucesso!', 'Usuário atualizado com sucesso', 'success').then(function() {
      window.location = "/perfil";
    })
  }).fail(function() {
    Swal.fire('Erro!', 'Erro ao atualizar o usuário', 'error');
  })
}

function atualizarSenha(event) {
  event.preventDefault();

  if ($('#nova-senha').val() != $('#confirmar-senha').val()) {
    Swal.fire('Erro!', 'As senha não coincidem', 'warning');
    return;
  }

  $.ajax({
    url: "atualizar-senha",
    method: "POST",
    data: {
      atual : $('#senha-atual').val(),
      nova : $('#nova-senha').val(),
    }
  }).done(function() {
    Swal.fire('Sucesso!', 'Senha atualizada com sucesso!', 'success').then(function() {
      window.location = "/perfil";
    })
  }).fail(function() {
    Swal.fire('Erro!', 'Erro ao atualizar a senha', 'error');
  })
}

function deletarUsuario(event) {
  event.preventDefault();
  Swal.fire({
    title: 'Atenção!',
    text: 'Tem certeza que deseja apagar a sua conta? Essa é uma ação irreversível!',
    showCancelButton: true,
    cancelButtonText: 'Cancelar',
    icon: 'warning',
  }).then(function(confirmacao) {
    if (confirmacao.value) {
      $.ajax({
        url: "/deletar-usuario",
        method: "DELETE",
      }).done(function() {
        Swal.fire('Sucesso!', 'A conta foi excluída permanentemente com sucesso!', 'success').then(function() {
          window.location = "/logout";
        })
      }).fail(function() {
        Swal.fire('Erro!', 'Erro ao excluir a conta!', 'error');
      })
    }
  })
}
