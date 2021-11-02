<template>
  <div id="app">
    <input type="text" v-model="text" @keydown.enter="wsHandler">
    <button @click="wsHandler" >click connect</button>
    <img :src="img" alt="">
    <!-- <p
      v-for="(item, index) in list"
      :key="index"
    >
      {{ item }}
    </p> -->
  </div>
</template>

<script>

export default {
  name: 'App',
  data() {
    return {
      text: '',
      list: [],
      img: '',
      ws: null
    };
  },
  created() {
    fetch('/api/test', {
      "Content-Type": 'application/json'
    }).then(res => {
      console.log('res.data :>> ', res);
    });
    this.ws = new WebSocket("ws://localhost:8000/ws");

    this.ws.onopen = () => {
      console.log('is connected');
    };
    this.ws.onmessage = (msg) => {
      console.log('msg :>> ', msg);
      console.log('msg.data :>> ', msg.data);
      this.img = msg.data;
      // this.list.push(msg.data)
    }
    this.ws.onclose = () => {
      console.log('close connect');
    }
    this.ws.onerror = (err) => {
      console.log('err :>> ', err);
    }
  },
  methods: {
    wsHandler() {
      this.ws.send(this.text);
      this.text = '';
    }
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
