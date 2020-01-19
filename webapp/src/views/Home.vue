<template>
  <q-page class="flex flex-center app-page">
    <HelloWorld msg="Welcome, vue.js "/>
  </q-page>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue';

export default {
  name: 'home',
  components: {
    HelloWorld,
  },
};

</script>

<style>
.app-page {
  padding: 16px 46px;
}
</style>
