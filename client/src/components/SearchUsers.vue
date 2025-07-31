<template>
    <div>
        <div class="custom-select-wrapper">
        <input v-model="searchWord" type="text" class="custom-input" @blur="hideDropDown" />

        <div v-if="isShowDropDown" class="custom-select-dropdown">
            <div
            v-for="user in searchResults"
            :key="user.id"
            class="custom-option"
            @mousedown.prevent="selectUser(user)"
            >
            {{ user.username }}
        </div>
  </div>
</div>
        <div>
            <h2>{{ selectedUser?.username }}</h2>
            <h3>{{ selectedUser?.pet.name }}</h3>
            <h3>{{ selectedUser?.pet.animal }}</h3>
            <h3>{{ selectedUser?.pet.favoriteFood }}</h3>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';

type Pet = {
    name: string;
    animal: string;
    favoriteFood: string;
}

type UserWithPet = {
    id: number;
    username: string;
    pet: Pet;
}

interface Props {
    SERVER_URL: string
}

const props = defineProps<Props>()

const searchWord = ref<string>("")
const searchResults = ref<UserWithPet[]>()
const controller = ref<AbortController>()
const selectedUser = ref<UserWithPet | null>()
const isShowDropDown = ref<boolean>(false)

const debounce = <T extends (...args: any[]) => any>(
  func: T,
  wait: number
): ((...args: Parameters<T>) => void) => {
  let timeoutId: ReturnType<typeof setTimeout> | undefined;

  return (...args: Parameters<T>) => {
    if (timeoutId) clearTimeout(timeoutId);

    timeoutId = setTimeout(() => {
      func(...args);
    }, wait);
  };
};

const selectUser = (user: UserWithPet) => {
    selectedUser.value = user
    searchResults.value = []
    isShowDropDown.value = false
}

const hideDropDown = () => {
    isShowDropDown.value = false
}

const search = async () => {
    if (controller.value){
        controller.value.abort()
    }
    try {
        selectedUser.value = null
        controller.value = new AbortController()
        const params = new URLSearchParams({
            userName: searchWord.value
        })
        const res = await fetch(`${props.SERVER_URL}/api/v2/search_users?${params.toString()}`, {
            signal: controller.value.signal
        })
        const data = await res.json()
        searchResults.value = data
        if (searchResults.value?.length) {
            isShowDropDown.value = true
        }
    } catch (error) {
        console.log(error)
    }
}

const debouncedSearch = debounce(search, 300)

watch(searchWord, () => {
    if(searchWord.value && searchWord.value.length > 1){
        debouncedSearch()
    } else {
        searchResults.value = []
    }
})

</script>

<style scoped>
.custom-select-wrapper {
  position: relative;
  width: 200px;
}

.custom-input {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.custom-select-dropdown {
  position: absolute;
  width: 100%;
  border: 1px solid #ccc;
  border-top: none;
  background-color: white;
  max-height: 200px;
  overflow-y: auto;
  z-index: 1000;
  box-shadow: 0px 2px 6px rgba(0, 0, 0, 0.2);
}

.custom-option {
  padding: 8px;
  cursor: pointer;
}

.custom-option:hover {
  background-color: #f0f0f0;
}
</style>