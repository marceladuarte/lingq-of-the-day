<template>
  <div>
    <div v-show="card.term">
      <div class="form-group">
        <div class="card-block">
          <div>
            <hr>
            <p>{{card.term}}</p>
            <div v-show="showHint">
              <p v-for="hint in card.hints" :key="hint.text" style="text-align:center">
                <i>({{hint.text}})</i>
              </p>
            </div>
          </div>
        </div> 
      </div>
      <div class="form-group">
        <div v-show="card.term && !showHint">
          <img title="Show meaning" alt="Show meaning" src="./../assets/eye.svg" height="30" @click="showHints()">
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Card',
  props: {
    card: Object
  },
  watch: {
    card: function(){
      this.showHint =  false
    }
  },
  data() {
    return{
      showHint: false
    }
  },

  methods: {
    showHints(){
      this.showHint = !this.showHint
    }
  }
}
</script>

<style scoped>
.card-block{
  margin: 0 auto; 
  width: 350px;
}
</style>
