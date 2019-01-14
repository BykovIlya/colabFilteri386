/* eslint-disable */
<template>
  <div id="app">
    <b-navbar toggleable="md" type="dark" variant="dark">
      <b-navbar-toggle target="nav_collapse"></b-navbar-toggle>
      <b-collapse is-nav id="nav_collapse">
        <b-navbar-nav>
          <b-nav-item :to = "{ name: 'Home'}" exact>Главная</b-nav-item>
          <b-nav-item :to = "{ name: 'Users'}">Покупатели</b-nav-item>
          <b-nav-item :to = "{ name: 'Products'}">Продукты</b-nav-item>
          <b-nav-item :to = "{ name: 'Recommendations'}">Персональные рекомендации</b-nav-item>
          <b-nav-item :to = "{ name: 'NeuralNetwork'}">Рекомендация для нового покупателя</b-nav-item>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /*margin-top: 60px;*/
}
</style>
