<script setup>
import { onMounted } from 'vue'
import { useMenusStore } from '../store/modules/menus'
import NavBarItem from "./NavBarItem.vue";

const menus = useMenusStore();
onMounted(() => menus.getAdminMenu());

</script>
<template>
    <div class="navbar-nav flex-row d-lg-none">
        <div class="nav-item d-none d-lg-flex me-3">
            <div class="btn-list">
                <a
                   href="https://github.com/tabler/tabler"
                   class="btn"
                   rel="noreferrer">
                    <svg
                         xmlns="http://www.w3.org/2000/svg"
                         class="icon"
                         width="24"
                         height="24"
                         viewBox="0 0 24 24"
                         strokeWidth="2"
                         stroke="currentColor"
                         fill="none"
                         strokeLinecap="round"
                         strokeLinejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                        <path
                              d="M9 19c-4.3 1.4 -4.3 -2.5 -6 -3m12 5v-3.5c0 -1 .1 -1.4 -.5 -2c2.8 -.3 5.5 -1.4 5.5 -6a4.6 4.6 0 0 0 -1.3 -3.2a4.2 4.2 0 0 0 -.1 -3.2s-1.1 -.3 -3.5 1.3a12.3 12.3 0 0 0 -6.2 0c-2.4 -1.6 -3.5 -1.3 -3.5 -1.3a4.2 4.2 0 0 0 -.1 3.2a4.6 4.6 0 0 0 -1.3 3.2c0 4.6 2.7 5.7 5.5 6c-.6 .6 -.6 1.2 -.5 2v3.5" />
                    </svg>
                    Source code
                </a>
                <a
                   href="https://github.com/sponsors/codecalm"
                   class="btn"
                   rel="noreferrer">
                    <svg
                         xmlns="http://www.w3.org/2000/svg"
                         class="icon text-pink"
                         width="24"
                         height="24"
                         viewBox="0 0 24 24"
                         strokeWidth="2"
                         stroke="currentColor"
                         fill="none"
                         strokeLinecap="round"
                         strokeLinejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                        <path d="M19.5 12.572l-7.5 7.428l-7.5 -7.428m0 0a5 5 0 1 1 7.5 -6.566a5 5 0 1 1 7.5 6.572" />
                    </svg>
                    Sponsor
                </a>
            </div>
        </div>

        <div class="d-none d-lg-flex">
            <a
               href="?theme=dark"
               class="nav-link px-0 hide-theme-dark"
               title="Enable dark mode"
               data-bs-toggle="tooltip"
               data-bs-placement="bottom">
                <svg
                     xmlns="http://www.w3.org/2000/svg"
                     class="icon"
                     width="24"
                     height="24"
                     viewBox="0 0 24 24"
                     strokeWidth="2"
                     stroke="currentColor"
                     fill="none"
                     strokeLinecap="round"
                     strokeLinejoin="round">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                    <path d="M12 3c.132 0 .263 0 .393 0a7.5 7.5 0 0 0 7.92 12.446a9 9 0 1 1 -8.313 -12.454z" />
                </svg>
            </a>

            <a
               href="?theme=light"
               class="nav-link px-0 hide-theme-light"
               title="Enable light mode"
               data-bs-toggle="tooltip"
               data-bs-placement="bottom">
                <svg
                     xmlns="http://www.w3.org/2000/svg"
                     class="icon"
                     width="24"
                     height="24"
                     viewBox="0 0 24 24"
                     strokeWidth="2"
                     stroke="currentColor"
                     fill="none"
                     strokeLinecap="round"
                     strokeLinejoin="round">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                    <circle cx="12" cy="12" r="4" />
                    <path
                          d="M3 12h1m8 -9v1m8 8h1m-9 8v1m-6.4 -15.4l.7 .7m12.1 -.7l-.7 .7m0 11.4l.7 .7m-12.1 -.7l-.7 .7" />
                </svg>
            </a>

            <div class="nav-item dropdown d-none d-md-flex me-3">
                <a
                   href="#"
                   class="nav-link px-0"
                   data-bs-toggle="dropdown"
                   aria-label="Show notifications">
                    <svg
                         xmlns="http://www.w3.org/2000/svg"
                         class="icon"
                         width="24"
                         height="24"
                         viewBox="0 0 24 24"
                         strokeWidth="2"
                         stroke="currentColor"
                         fill="none"
                         strokeLinecap="round"
                         strokeLinejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                        <path
                              d="M10 5a2 2 0 0 1 4 0a7 7 0 0 1 4 6v3a4 4 0 0 0 2 3h-16a4 4 0 0 0 2 -3v-3a7 7 0 0 1 4 -6" />
                        <path d="M9 17v1a3 3 0 0 0 6 0v-1" />
                    </svg>

                    <span class="badge bg-red"></span>
                </a>
                <div class="dropdown-menu dropdown-menu-arrow dropdown-menu-end dropdown-menu-card">
                    <div class="card">
                        <div class="card-header">
                            <h3 class="card-title">Last updates</h3>
                        </div>
                        <div class="list-group list-group-flush list-group-hoverable">
                            <div class="list-group-item">
                                <div class="row align-items-center">
                                    <div class="col-auto">
                                        <span class="status-dot status-dot-animated bg-red d-block"></span>
                                    </div>
                                    <div class="col text-truncate">
                                        <a href="#" class="text-body d-block">
                                            Example 1
                                        </a>
                                        <div class="d-block text-muted text-truncate mt-n1">
                                            Change deprecated html tags to text decoration classes
                                            (#29604)
                                        </div>
                                    </div>
                                    <div class="col-auto">
                                        <a href="#" class="list-group-item-actions">
                                            <svg
                                                 xmlns="http://www.w3.org/2000/svg"
                                                 class="icon text-muted"
                                                 width="24"
                                                 height="24"
                                                 viewBox="0 0 24 24"
                                                 strokeWidth="2"
                                                 stroke="currentColor"
                                                 fill="none"
                                                 strokeLinecap="round"
                                                 strokeLinejoin="round">
                                                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                                <path
                                                      d="M12 17.75l-6.172 3.245l1.179 -6.873l-5 -4.867l6.9 -1l3.086 -6.253l3.086 6.253l6.9 1l-5 4.867l1.179 6.873z" />
                                            </svg>
                                        </a>
                                    </div>
                                </div>
                            </div>
                            <div class="list-group-item">
                                <div class="row align-items-center">
                                    <div class="col-auto">
                                        <span class="status-dot d-block"></span>
                                    </div>
                                    <div class="col text-truncate">
                                        <a href="#" class="text-body d-block">
                                            Example 2
                                        </a>
                                        <div class="d-block text-muted text-truncate mt-n1">
                                            justify-content:between ⇒
                                            justify-content:space-between (#29734)
                                        </div>
                                    </div>
                                    <div class="col-auto">
                                        <a href="#" class="list-group-item-actions show">
                                            <svg
                                                 xmlns="http://www.w3.org/2000/svg"
                                                 class="icon text-yellow"
                                                 width="24"
                                                 height="24"
                                                 viewBox="0 0 24 24"
                                                 strokeWidth="2"
                                                 stroke="currentColor"
                                                 fill="none"
                                                 strokeLinecap="round"
                                                 strokeLinejoin="round">
                                                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                                <path
                                                      d="M12 17.75l-6.172 3.245l1.179 -6.873l-5 -4.867l6.9 -1l3.086 -6.253l3.086 6.253l6.9 1l-5 4.867l1.179 6.873z" />
                                            </svg>
                                        </a>
                                    </div>
                                </div>
                            </div>
                            <div class="list-group-item">
                                <div class="row align-items-center">
                                    <div class="col-auto">
                                        <span class="status-dot d-block"></span>
                                    </div>
                                    <div class="col text-truncate">
                                        <a href="#" class="text-body d-block">
                                            Example 3
                                        </a>
                                        <div class="d-block text-muted text-truncate mt-n1">
                                            Update change-version.js (#29736)
                                        </div>
                                    </div>
                                    <div class="col-auto">
                                        <a href="#" class="list-group-item-actions">
                                            <svg
                                                 xmlns="http://www.w3.org/2000/svg"
                                                 class="icon text-muted"
                                                 width="24"
                                                 height="24"
                                                 viewBox="0 0 24 24"
                                                 strokeWidth="2"
                                                 stroke="currentColor"
                                                 fill="none"
                                                 strokeLinecap="round"
                                                 strokeLinejoin="round">
                                                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                                <path
                                                      d="M12 17.75l-6.172 3.245l1.179 -6.873l-5 -4.867l6.9 -1l3.086 -6.253l3.086 6.253l6.9 1l-5 4.867l1.179 6.873z" />
                                            </svg>
                                        </a>
                                    </div>
                                </div>
                            </div>
                            <div class="list-group-item">
                                <div class="row align-items-center">
                                    <div class="col-auto">
                                        <span class="status-dot status-dot-animated bg-green d-block"></span>
                                    </div>
                                    <div class="col text-truncate">
                                        <a href="#" class="text-body d-block">
                                            Example 4
                                        </a>
                                        <div class="d-block text-muted text-truncate mt-n1">
                                            Regenerate package-lock.json (#29730)
                                        </div>
                                    </div>
                                    <div class="col-auto">
                                        <a href="#" class="list-group-item-actions">
                                            <svg
                                                 xmlns="http://www.w3.org/2000/svg"
                                                 class="icon text-muted"
                                                 width="24"
                                                 height="24"
                                                 viewBox="0 0 24 24"
                                                 strokeWidth="2"
                                                 stroke="currentColor"
                                                 fill="none"
                                                 strokeLinecap="round"
                                                 strokeLinejoin="round">
                                                <path stroke="none" d="M0 0h24v24H0z" fill="none" />
                                                <path
                                                      d="M12 17.75l-6.172 3.245l1.179 -6.873l-5 -4.867l6.9 -1l3.086 -6.253l3.086 6.253l6.9 1l-5 4.867l1.179 6.873z" />
                                            </svg>
                                        </a>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="nav-item dropdown">
            <a
               href="#"
               class="nav-link d-flex lh-1 text-reset p-0"
               data-bs-toggle="dropdown"
               aria-label="Open user menu">
                <span
                      class="avatar avatar-sm"
                      :style="{ 'background-image': 'url(./static/avatars/000m.jpg)' }"></span>

                <div class="d-none d-xl-block ps-2">
                    <div>Paweł Kuna</div>
                    <div class="mt-1 small text-muted">UI Designer</div>
                </div>
            </a>
            <div class="dropdown-menu dropdown-menu-end dropdown-menu-arrow">
                <a href="#" class="dropdown-item">
                    Set status
                </a>

                <a href="#" class="dropdown-item">
                    Profile & account
                </a>

                <a href="#" class="dropdown-item">
                    Feedback
                </a>

                <div class="dropdown-divider"></div>

                <a href="#" class="dropdown-item">
                    Settings
                </a>

                <a href="#" class="dropdown-item">
                    Logout
                </a>
            </div>
        </div>
    </div>
    <div id="navbar-menu" class="collapse navbar-collapse">
        <ul class="navbar-nav pt-lg-3">
            <template v-for="navItem in menus.admin?.items" :key="navItem.id">
                <NavBarItem :item="navItem" />
            </template>
        </ul>
    </div>
</template>


