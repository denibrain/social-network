import "../../Layout/index.css"
import "../../Components/SignUpForm/index.css"
import * as $ from 'jquery'

$(".sign-in").submit(function () {
    $.post("/signup", {
        email: $("#input-email").val(),
        password: $("#input-password").val(),

        name: $("#input-name").val(),
        surname: $("#input-surname").val(),

        city: $("#input-city").val(),
        interests: $("#input-interests").val(),

        age: $("#input-age").val(),
        sex: $("#input-sex").val(),

    }, function (response) {
        document.location = response.returnTo;
    }, "json")
    .fail(function(xhr) {
        if (xhr.status == 400 || xhr.status == 409) {
            if (xhr.responseJSON.field) {
                let fieldName = $("label[for$=" + xhr.responseJSON.field + "]").text();
                error(fieldName + ": " +xhr.responseJSON.error);
            }
            else
            {
                error(xhr.responseJSON.error);
            }
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