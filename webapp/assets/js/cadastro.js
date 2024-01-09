$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
  event.preventDefault();

  if ($('#senha').val() != $('#confirmarSenha').val()) {
    Swal.fire('Erro!', 'As senhas não conferem!', 'error');
    return;
  }

  $.ajax({
    url: "/usuarios",
    method: "POST",
    data: {
      nome: $('#nome').val(),
      email: $('#email').val(),
      nick: $('#nick').val(),
      senha: $('#senha').val()
    }
  }).done(function() {
    Swal.fire('Sucesso!', 'Usuário cadastrado!', 'success').then(function() {
      $.ajax({
        url: "/login",
        method: "POST",
        data: {
          email: $('#email').val(),
          senha: $('#senha').val()
        }
      }).done(function() {
        window.location = "/home";
      }).fail(function() {
        Swal.fire('Erro!', 'Erro ao efetuar login!', 'error');
      })
    });
  }).fail(function(erro) {
    Swal.fire('Erro!', 'Erro ao cadastrar usuário: ' + erro?.responseJSON?.erro, 'error');
  })
}