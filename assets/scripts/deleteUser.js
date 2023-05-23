$("#deleteUser").on("click", deleteUser);

function deleteUser(event) {
    event.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir sua conta? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning",
    }).then(function (confirmation) {
        if (!confirmation.isConfirmed) return;

        handleUserDelete();
    });
}

function handleUserDelete() {
    $.ajax({
        url: "/api/userDelete",
        method: "DELETE",
    })
        .done(function () {
            Swal.fire("Sucesso", "Conta excluida com sucesso!", "success").then(
                function () {
                    window.location = "/logout";
                }
            );
        })
        .fail(function () {
            Swal.fire("Ops...", "Ocorreu um erro ao excluir sua conta!", "error");
        });
}
