$("#create-user-form").on("submit", createUser);

function createUser(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    if (data.password != data.confirmPassword) {
        alert("Senhas não coincidem!");
        return;
    }

    $.ajax({
        url: "api/user",
        method: "POST",
        data,
    })
        .done(function (data) {
            alert(JSON.stringify(data));
            alert("Usuário Cadastrado com sucesso!");
        })
        .fail(function (data) {
            console.log(JSON.stringify(data));
        });
}

function getFormFieldsValues() {
    const name = $("#name").val();
    const email = $("#email").val();
    const nickname = $("#nickName").val();
    const password = $("#password").val();
    const confirmPassword = $("#confirmPassword").val();

    return { name, email, nickname, password, confirmPassword };
}
