$('#formulario-atualiza-cadastro').on('submit', atualizarUser);
$('#formulario-atualizar-senha-cadastro').on('submit', atualizarSenhaUser);
$('.btn-trash').on('click', deletarPost);

function deletarPost(evento) {
    evento.preventDefault();
    const idPost = $(this).closest('.card').data('publicacao-id');

    Swal.fire({
        title: "Atenção!",
        text: "tem certeza que deseja excluir essa Publicação?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirmacao) => {
        if (!confirmacao.isConfirmed) return;

        $.ajax({
            url: `/publicacoes/${idPost}`,
            type: 'DELETE',
            dataType: 'json',
        }).done(function (data) {
            Swal.fire({
                title: "Publicacao deletada com sucesso!",
                icon: "success"
            }).then(() => {
                window.location.reload();
            });
        }).fail(function (data) {
            console.log(data.responseJSON);
            Swal.fire({
                title: "Falha ao deletar publicação!",
                icon: "error"
            });
        });

    });
}

function atualizarUser(evento) {
    evento.preventDefault();
    var id = $('#userID').val();
    Swal.fire({
        title: "Atenção!",
        text: "Você que realmente Atualizar seus Dados?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirmacao) => {
        if (!confirmacao.isConfirmed) return;

        $.ajax({
            url: `/usuario/${id}`,
            type: 'PUT',
            dataType: 'json',
            data: {
                nome: $('#nome').val(),
                nick: $('#nick').val(),
                email: $('#email').val(),
            }
        }).done(function (data) {
            Swal.fire({
                title: "Atualizado com sucesso!",
                icon: "success"
            }).then(() => {
                window.location.reload();
            });
        }).fail(function (data) {
            console.log(data.responseJSON);
            Swal.fire({
                title: "Falha ao atualizar usuario!",
                text: "Nick ou email ja em uso",
                icon: "error"
            });
        });

    });
}

function atualizarSenhaUser(evento) {
    evento.preventDefault();
    var id = $('#userID').val();
    if ($('#senha').val() != $('#ConfirmarSenha').val()) {
        Swal.fire({
            title: "Atenção!",
            text: "Senha nao coencidem!",
            icon: "error"
        })
        return
    }

    Swal.fire({
        title: "Atenção!",
        text: "Você que realmente Atualizar a senha",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then((confirmacao) => {
        if (!confirmacao.isConfirmed) return;

        $.ajax({
            url: `/usuario/${id}/atualizar-pass`,
            type: 'PUT',
            dataType: 'json',
            data: {
                atual: $('#atual').val(),
                senha: $('#senha').val(),
            }
        }).done(function (data) {
            Swal.fire({
                title: "Senha atualizada com sucesso!",
                icon: "success"
            }).then(() => {
                window.location.reload();
            });
        }).fail(function (data) {
            console.log(data.responseJSON);
            Swal.fire({
                title: "Falha ao atualizar senha!",
                icon: "error"
            });
        });

    });
}