// script.js
document.addEventListener('DOMContentLoaded', function() {
    const registerForm = document.getElementById('register-form');
    const loginFormBox = document.querySelector('.form-box.login');
    const registerFormBox = document.querySelector('.form-box.register');

    // Menampilkan form login saat tombol Join Us diklik
    document.getElementById('join-us-btn').onclick = function() {
        registerForm.style.display = 'block'; // Tampilkan form
        loginFormBox.style.display = 'block'; // Tampilkan form login secara default
        registerFormBox.style.display = 'none'; // Sembunyikan form register saat pertama kali terbuka
    }

    // Menutup form saat tombol Close diklik
    document.getElementById('close-form').onclick = function() {
        registerForm.style.display = 'none'; // Sembunyikan form utama
    }

    // Menangani klik pada link "Register" di form login
    document.querySelector('.register-link').onclick = function(e) {
        e.preventDefault();
        loginFormBox.style.display = 'none'; // Sembunyikan form login
        registerFormBox.style.display = 'block'; // Tampilkan form register
    }

    // Menangani klik pada link "Login" di form register
    document.querySelector('.login-link').onclick = function(e) {
        e.preventDefault();
        registerFormBox.style.display = 'none'; // Sembunyikan form register
        loginFormBox.style.display = 'block'; // Tampilkan form login
    }

    // Menangani pengiriman form login
    document.getElementById('login-form').onsubmit = function(event) {
        event.preventDefault(); 
    }
});
