<template>
    <div class="container">
        <form class="form-horizontal"  @submit.prevent="attemptLogin({ username, password, country })">

            <div class="row">
                <div class="col-md-6 mx-auto">
                    <div class="alert alert-danger" v-if="showLoginError">
                        Invalid username or password
                    </div>
                </div>
            </div>

            <div class="row">
                <div class="col-md-6 mx-auto">
                    <div class="form-group has-danger">
                        <div class="input-group mb-2 mr-sm-2 mb-sm-0">
                            <div class="input-group-addon"><i class="fa fa-at"></i></div>
                            <input name="username" class="form-control" v-model="username" placeholder="Username" required autofocus>
                        </div>
                    </div>
                </div>
            </div>

            <div class="row">
                <div class="col-md-3"></div>
                <div class="col-md-6">
                    <div class="form-group">
                        <div class="input-group mb-2 mr-sm-2 mb-sm-0">
                            <div class="input-group-addon"><i class="fa fa-key"></i></div>
                            <input type="password" class="form-control" v-model="password" placeholder="Password" required>
                        </div>
                    </div>
                </div>
            </div>

            <div class="row">
                <div class="col-md-3"></div>
                <div class="col-md-6">
                    <div class="form-group">
                        <div class="input-group mb-2 mr-sm-2 mb-sm-0">
                            <div class="input-group-addon"><i class="fa fa-key"></i></div>
                            <select class="form-control" v-model="country" required>
                                <option value="us" selected>US</option>
                                <option value="de">DE</option>
                                <option value="uk">UK</option>
                                <option value="be">BE</option>
                                <option value="ca">CA</option>
                                <option value="ch">CH</option>
                                <option value="au">AU</option>
                            </select>
                        </div>
                    </div>
                </div>
            </div>

            <div class="row">
                <div class="col-md-3"></div>
                <div class="col-md-6">
                    <button type="submit" class="btn btn-success btn-block"><i class="fa fa-sign-in"></i> Login</button>
                </div>
            </div>

        </form>
    </div>
</template>

<script>
import { login } from '@/api/auth'
import { setAccessToken, setUserRole, setUserId, setUserName } from '@/api'

export default {
    data() {
        return {
            username: '',
            password: '',
            showLoginError: false,
            country : 'us'
        }
    },

    methods: {
        attemptLogin({ username, password, country }) {
            const redirectTo = this.$route.query.redirect ? this.$route.query.redirect : '/';

            login(username, password, country)
                .then(({ data }) => {
                    setAccessToken(data.access_token)
                    setUserId(data.user_data.id)
                    setUserName(data.user_data.username)
                    setUserRole(data.user_data.metadata.roles)
                    this.$router.push(redirectTo)
                })
                // .catch(() => this.showLoginError = true)
        }
    }
}
</script>

<style scoped>
.container {
    margin-top: 4rem;
}
</style>
