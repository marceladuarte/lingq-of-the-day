<template>
  <div> 
    <div class="error-block">
      <hr v-show="msg">
      <div>
        <p>{{msg}}</p>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Error',
  props: {
    msg: String
  }
}
</script>

<style scoped>
.error-block{
  margin: 0 auto; 
  width: 350px;
}
</style>