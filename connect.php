<?php
$servername = "localhost";
$username = "form"; // ganti dengan username MySQL Anda
$password = ""; // ganti dengan password MySQL Anda
$dbname = "form_database"; // ganti dengan nama database Anda

$conn = new mysqli($servername, $username, $password, $dbname);

if ($conn->connect_error) {
    die("Koneksi gagal: " . $conn->connect_error);
}
?>
