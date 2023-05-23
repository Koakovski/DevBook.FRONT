$("#form-update-password").on("submit", updateUserPassword);

function updateUserPassword(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    if (data.newPassword !== data.confirmationPassword) {
        Swal.fire("Ops...!", "As senhas não coincidem!", "warning");
        return;
    }

    $.ajax({
        url: "/api/updatePassword",
        method: "POST",
        data,
    })
        .done(function () {
            Swal.fire("Sucesso!", "Senha Atualizada com sucesso!", "success").then(() => {
                window.location = "/profile";
            });
        })
        .fail(function () {
            Swal.fire("Ops...", "Falha ao atualizar a senha!", "error");
        });
}

function getFormFieldsValues() {
    const currentPassword = $("#currentPassword").val();
    const newPassword = $("#newPassword").val();
    const confirmationPassword = $("#confirmationPassword").val();

    return { currentPassword, newPassword, confirmationPassword };
}
