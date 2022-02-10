<script setup>
import { ref } from "vue";
import axios from "axios";
const email = ref("");
const isEmpty = ref(false);
const emailSent = ref(false);

function verifySubmit() {
  console.log(email);
  if (email.value == "") {
    isEmpty.value = true;
    return false;
  }

  axios
    .post("/resetemail", { email: email.value.toLowerCase() })
    .then(function (response) {
      // Redirect to login page maybe?
      console.log(response);
      emailSent.value = true;
    })
    .catch(function (error) {
      console.log(error);
    });
}
</script>

<template>
  <div
    v-if="!emailSent"
    class="container mt-3 reset-form d-flex justify-content-center align-items-center min-vh-100"
  >
    <form @submit.prevent="verifySubmit">
      <div class="mb-3 mt-3">
        <label for="email">Email:</label>
        <input
          type="email"
          class="form-control"
          id="email"
          placeholder="Email address"
          name="email"
          v-model="email"
        />
      </div>
      <button type="submit" class="btn btn-lg">Submit</button>
    </form>
  </div>
  <div
    v-else
    class="container mt3 success d-flex flex-column justify-content-center align-items-center min-vh-100"
  >
    <h1>An email has been sent!</h1>
  </div>
</template>

<style>
@import "./assets/base.css";

.btn-lg {
  background-color: #dcf900;
  color: black;
  text-align: center;
  border-color: rgb(219, 238, 79);
  font-weight: 700;
  width: 100%;
  vertical-align: middle;
}
</style>
