<script setup lang="ts">

import { ref } from "vue"
import BaseIcon from "./BaseIcon.vue"
import { useRouter/*, useRoute */ } from 'vue-router'

export interface IMenuItem {
    id: string
    label: string
    path: string
    children: IMenuItem[]
}
export interface Props {
    item: IMenuItem
}
defineProps<Props>();

const router = useRouter()
//const route = useRoute()
const expanded = ref(false)

const isActive = ref(false)

const toggle = () => {
    expanded.value = !expanded.value
}

function gotoRoute(path: string) {
    router.push({
        name: path
    })
}

const slugify = (s: string) => s.toLowerCase().replace(/\s+/g, '-').replace(/[^\w-]+/g, '');


</script>
<template>
    <li
        class="nav-item"
        :class="{ dropdown: item.children?.length > 0, active: isActive }">
        <a
           :class="'nav-link' + (item.children?.length > 0 ? ' dropdown-toggle' : '')"
           data-bs-toggle="dropdown"
           data-bs-auto-close="false"
           role="button"
           :aria-expanded="expanded"
           @click="item.children?.length > 0 ? toggle : gotoRoute(item.path)">
            <span class="nav-link-icon d-md-none d-lg-inline-block">
                <BaseIcon :name="slugify(item.label)" />
            </span>
            <span class="nav-link-title">{{ item.label }}</span>
        </a>
        <div v-if="item.children?.length > 0" :class="'dropdown-menu' + (expanded ? ' show' : '')">
            <div class="dropdown-menu-columns">
                <div class="dropdown-menu-column">
                    <a v-for="cItem in item.children" :key="cItem.id" class="dropdown-item">
                        {{ cItem.label }}
                    </a>
                </div>
            </div>
        </div>
    </li>
</template>