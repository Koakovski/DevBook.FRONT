$("#form-update-user").on("submit", updateUser);

function updateUser(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    $.ajax({
        url: "/api/updateUser",
        method: "PUT",
        data,
    })
        .done(function () {
            Swal.fire("Sucesso!", "Usuário atualizado com sucesso!", "success").then(
                () => {
                    window.location = "/profile";
                }
            );
        })
        .fail(function () {
            Swal.fire("Ops..", "Falha ao atualizar a usuário!", "error");
        });
}

function getFormFieldsValues() {
    const name = $("#name").val();
    const email = $("#email").val();
    const nickName = $("#nickName").val();

    return { name, email, nickName };
}
