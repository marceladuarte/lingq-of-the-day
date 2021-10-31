<template>
  <div>
    <form @submit.prevent>
      <div v-show="languages">
        <div class="form-group">
          <select name="languages" id="languages" v-model="language" >
            <option v-for="lang in languages" :key="lang.code" :value="lang.code">{{ lang.title }}</option>
          </select>
        </div>
        <div class="form-group">
          <button @click="getCard()">Select LingQ!</button>
        </div>
        <Card :card="card"/>
      </div>
      <Error :msg="errorMsg"/>
    </form>
  </div>
</template>

<script>

import axios from 'axios'
import Card from './Card.vue'
import Error from './Error.vue'

export default {
  components: { Card, Error },
  name: 'Language',
  data() {
    return{
      languages: [],
      language: 'en',
      card: {},
      errorMsg: null
    }
  },

  methods: {
    sortLanguages(){
      this.languages.sort((a,b) => {
        if(a.title < b.title){
          return -1
        }
        if(a.title > b.title){
          return 1;
        }
        return 0;
      })
    },

    async getCard(){
      this.errorMsg = null
      try{
        const resp = await axios.get(`${process.env.VUE_APP_SERVER_URL}/api/card?lang=${this.language}`, {headers: {'Origin': process.env.VUE_APP_SERVER_URL}})
        this.card = resp.data
      }catch(error){
        console.error(error.response)
        this.card = {}
        if(error.response.status == 400) {
          this.errorMsg = "Couldn't find a LingQ for the selected language."
        }else{
          this.errorMsg = "Couldn't retrieve a LingQ. Try again later."
        }
      }
    }
  },

  async mounted() {
    try{
      const resp = await axios.get(`${process.env.VUE_APP_SERVER_URL}/api/languages`, {headers: {'Origin': process.env.VUE_APP_SERVER_URL}})
      this.languages = resp.data
      this.sortLanguages()
    }catch(error){
      console.error(error.response)
      this.languages = null
      this.errorMsg = "Couldn't retrieve languages from LingQ. Try again later."
    }
  }
}
</script>

<style scoped>
.form-group{
  margin-bottom: 30px;
}
button{
  background-color: #2c3e50;
  border: none;
  color: white;
  padding: 15px 32px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
}
select{
  height: 48px;
  text-align: center;
  width: 160px;
}
</style>