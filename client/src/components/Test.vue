<template>
    <div>
        <h1>hello</h1>
        <button @click="test">Hey</button>
    </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';

const test = () => {
    console.log('test')
}

onMounted(() => {
    console.log('hi')
})
</script>

<style scoped>

</style>