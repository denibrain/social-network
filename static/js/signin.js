$(".sign-in").submit(function () {
    $.post("/signin", {
        username: $("#input-email").val(),
        password: $("#input-password").val(),
    }, function (response) {
        document.location = response.returnTo;
    }, "json")
    .fail(function(xhr) {
        if (xhr.status == 401) {
            error("Username or Password is invalid");
            return false;
        }
        error("Internal system error")
    });

    return false;
});

function error(message) {
    $(".message-box").html(
        "<div class=\"alert alert-danger\">" +
        message +
        "</div>"
    )
}