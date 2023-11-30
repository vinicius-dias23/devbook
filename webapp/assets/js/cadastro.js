$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
  event.preventDefault();

  if ($('#senha').val() != $('#confirmarSenha').val()) {
    alert("As senhas não conferem!");
    return
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
  })
}