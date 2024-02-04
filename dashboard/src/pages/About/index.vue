<template>
  <div class="absolute inset-x-0 bottom-0">
  </div>
</template>

<script>
import http from '@/api/http';

export default {
  name: "About",

  mounted() {
    http.get("/nodes/").then(response => {
      console.log(response.data)
    })
  }
}
</script>


<style scoped>

</style>