<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png">   -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <Test/>
    <!-- <Example1/> -->
    <!-- <Example2/> -->
    <!-- <Example3/> -->
  </div>
</template>

<script>
//import HelloWorld from './components/HelloWorld.vue'
import Test from './components/Test.vue'
//import Example1 from './components/Example1.vue'
//import Example2 from './components/Example2.vue'
//import Example3 from './components/Example3.vue'

export default {
  name: 'App',
  components: {
    //HelloWorld,
    Test
    //Example1
    //Example2
    //Example3
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
