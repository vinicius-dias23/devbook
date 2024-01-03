$('#nova-publicacao').on('submit', criarPublicacao);

function criarPublicacao(event) {
  event.preventDefault;

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