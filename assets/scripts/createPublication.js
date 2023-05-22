$("#create-publication-form").on("submit", CreatePublication);

function CreatePublication(event) {
    event.preventDefault();

    const data = getFormFieldsValues();

    $.ajax({
        url: "api/publication",
        method: "POST",
        data,
    })
        .done(function () {
            alert("Publicação criada com sucesso!");
        })
        .fail(function () {
            alert("Falha ao criar publicação!");
        });
}

function getFormFieldsValues() {
    const title = $("#title").val();
    const content = $("#content").val();

    return { title, content };
}
