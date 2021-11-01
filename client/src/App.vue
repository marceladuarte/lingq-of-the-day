<template>
  <img alt="Chat logo" src="./assets/chat.svg" height="200">
  <h1>LingQ of The Day</h1>
  <Language/>
</template>

<script>
import Language from './components/Language.vue'

export default {
  name: 'App',
  components: {
    Language
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
