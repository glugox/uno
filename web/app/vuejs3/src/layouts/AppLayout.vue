<script setup>
import NavBarLogo from "../components/NavBarLogo.vue"
import NavBar from "../components/NavBar.vue"
</script>
<template>
    <div className="page">
        <aside className="navbar navbar-vertical navbar-expand-lg navbar-dark">
            <div className="container-fluid">
                <button
                        className="navbar-toggler"
                        type="button"
                        data-bs-toggle="collapse"
                        data-bs-target="#navbar-menu">
                    <span className="navbar-toggler-icon"></span>
                </button>
                <NavBarLogo />
                <NavBar />

            </div>
        </aside>
        <div className="page-wrapper">
            <div className="container-xl">
                <slot />
            </div>
        </div>
    </div>
</template>
