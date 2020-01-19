<template>
  <div class="flex flex-center">
    <img alt="Quasar logo" src="../assets/logo.png">
    <h3>{{msg}}</h3>
  </div>
</template>

<style>
</style>

<script>
export default {
  props: ['msg'],
  name: 'HelloWorld',
};
</script>
