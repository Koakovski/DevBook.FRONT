$("#create-user-form").on("submit", createUser);

function createUser(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    if (data.password != data.confirmPassword) {
        Swal.fire("Ops..", "As senha não coincidem!", "error");
        return;
    }

    $.ajax({
        url: "/api/user",
        method: "POST",
        data,
    })
        .done(successHandler)
        .fail(function (jqXHR) {
            if (jqXHR.status === 201) {
                successHandler();
                return;
            }

            Swal.fire("Ops..", "Falha ao cadastror o usuário!", "error");
        });
}

function successHandler() {
    Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success").then(function () {
        window.location = "/login";
    });
}

function getFormFieldsValues() {
    const name = $("#name").val();
    const email = $("#email").val();
    const nickName = $("#nickName").val();
    const password = $("#password").val();
    const confirmPassword = $("#confirmPassword").val();

    return { name, email, nickName, password, confirmPassword };
}
