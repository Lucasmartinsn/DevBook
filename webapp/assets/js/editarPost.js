$('#editar-publicacao').on('submit', editarPost);

function editarPost(evento) {
    evento.preventDefault();
    $(this).prop('disabled', true)
    var id = $('#postID').val();

    Swal.fire({
        title: "Atenção!",
        text: "tem certeza que deseja atualizar essa Publicação?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirmacao) => {
        if (!confirmacao.isConfirmed) return;
    
        $.ajax({
            url: `/publicacoes/${id}`, // URL para onde enviar o POST
            type: 'PUT', // Método HTTP a ser utilizado (GET, POST, etc.)
            dataType: 'json', // Tipo de dado esperado de retorno
            data: {
                titulo: $('#titulo').val(),
                conteudo: $('#conteudo').val(),
            }
        }).done(function(data) {
            Swal.fire({
                title: "Publicação atualizada com Sucesso!",
                icon: "success"
            }).then(() => {
            history.back();
            });
        }).fail(function(data) {
            console.log(data);
            Swal.fire({
                title: "Falha ao atualizar publicação!",
                icon: "error"
            });
        }).always(function() {
            $('#editar-publicacao').prop('disabled', false)
        }) ;
        
    });
}