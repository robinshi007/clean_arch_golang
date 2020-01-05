<template>
  <q-page class="flex flex-center">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
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
