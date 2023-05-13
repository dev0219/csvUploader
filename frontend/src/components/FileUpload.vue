<template>
  <div class="file-upload">
    <input type="file" @change="onFileChange" />
    <div v-if="progress" class="progess-bar" :style="{'width': progress}">{{progress}}</div>
    <button @click="onUploadFile" class="upload-button" :disabled="!this.selectedFile">Upload file</button>
  </div>
</template>

<script>
import axios from "axios";


/* eslint-disable */
export default {
  data() {
    return {
      selectedFile: "",
      progress: 0
    };
  },
  created() {
    
  },
  methods: {
    onFileChange(e) {
      const selectedFile = e.target.files[0]; // accessing file
      this.selectedFile = selectedFile;
      this.progress = 0;
    },
    onUploadFile() {
      const formData = new FormData();
      formData.append("filename", this.selectedFile); // appending file

      axios
        .post("http://localhost:8000/uploadcsv", formData, {
          onUploadProgress: ProgressEvent => {
            let progress =
              Math.round((ProgressEvent.loaded / ProgressEvent.total) * 100) +
              "%";
            this.progress = progress;
          }
        })
        .then(res => {
          if (res.status == 200 && res.data.failedRows == null) {
            this.$toast("file uploaded successfully!", {
              position: "top-right",
              timeout: 6000,
              closeOnClick: true,
              pauseOnFocusLoss: true,
              pauseOnHover: true,
              draggable: true,
              draggablePercent: 0.6,
              showCloseButtonOnHover: false,
              hideProgressBar: true,
              closeButton: "button",
              icon: true,
              rtl: false
            });
          }         
        })
        .catch(err => {
          console.log(err);
          this.$toast("Please try to upload the file again.", {
              position: "top-right",
              timeout: 6000,
              closeOnClick: true,
              pauseOnFocusLoss: true,
              pauseOnHover: true,
              draggable: true,
              draggablePercent: 0.6,
              showCloseButtonOnHover: false,
              hideProgressBar: true,
              closeButton: "button",
              icon: true,
              rtl: false
            });
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.file-upload {
  box-shadow: 2px 2px 9px 2px #ccc;
  border-radius: 1rem;
  padding: 2rem;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  font-size: 1rem;
  width: 60%;
  margin: 0 auto;
}

input {
  font-size: 0.9rem;
}

input,
div,
button {
  margin-top: 2rem;
}

.progess-bar {
  height: 1rem;
  width: 0%;
  background-color: #151618;
  color: white;
  padding: 2px;
  font-weight: bold;
}

.upload-button {
  width: 7rem;
  padding: 0.5rem;
  background-color: #278be9;
  color: #fff;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  border-radius: 1rem;
}

.upload-button:disabled {
  background-color: #b3bcc4;
  cursor: no-drop;
}
</style>
