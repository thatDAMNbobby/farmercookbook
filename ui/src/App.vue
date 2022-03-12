            <template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/>
</template>

<script>
import Recipe from './components/Recipe.vue'

const rec = {
  id: 1,
  name: "food",
  author: "bobby",
  submitter: "bobby",
  course: "entree",
  tags: ["beef"],
  ingredients: [
    {
      name: "salt",
      quantity: {
        amount: 1,
        measure: "Tablespoon"
      }
    }
  ],
  canDouble: true
}
export default {
  name: 'App',
  Components: {
    Recipe: rec
  }
}
Recipe.data = function() {
  return {
    recipe: rec
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
