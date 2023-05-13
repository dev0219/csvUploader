<template>
  <div id="app">
    <FileUpload />
  </div>
</template>

<script>
import Vue from "vue";
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";
Vue.use(Toast, {
  transition: "Vue-Toastification__bounce",
  maxToasts: 20,
  newestOnTop: true
});
import FileUpload from "./components/FileUpload.vue";

export default {
  name: "app",
  components: {
    FileUpload
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #151618e0;
  margin: 60px auto;
  max-width: 1000px;
}
</style>
