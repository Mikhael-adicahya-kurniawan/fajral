document.addEventListener('DOMContentLoaded', function() {
    const registerForm = document.getElementById('register-form');
    const loginFormBox = document.querySelector('.form-box.login');
    const registerFormBox = document.querySelector('.form-box.register');

    if (registerForm && loginFormBox && registerFormBox) {
        document.getElementById('join-us-btn').onclick = function() {
            registerForm.style.display = 'flex';
            loginFormBox.style.display = 'block';
            registerFormBox.style.display = 'none';
        };

        document.getElementById('close-form').onclick = function() {
            registerForm.style.display = 'none';
        };

        document.querySelector('.register-link').onclick = function(e) {
            e.preventDefault();
            loginFormBox.style.display = 'none';
            registerFormBox.style.display = 'block';
        };

        document.querySelector('.login-link').onclick = function(e) {
            e.preventDefault();
            registerFormBox.style.display = 'none';
            loginFormBox.style.display = 'block';
        };
    }

    const loginForm = document.getElementById('login-form');
    if (loginForm) {
        loginForm.onsubmit = function(event) {
            event.preventDefault();
            // Tambahkan fungsi login di sini
        };
    }
});
