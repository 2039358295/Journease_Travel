import Vuex from 'vuex';


export default new Vuex.Store({
    state: {
        selectedTabIndex: 0,
        searchContent:'',
        usernickname:'',
        useravadar:'',
    },
    mutations: {
        setSelectedTabIndex(state, index) {
            state.selectedTabIndex = index;
        },
        setSearchContent(state, content){
            console.log(content)
            state.searchContent = content;
        },
        sethead(state,data){
            state.usernickname=data.usernickname;
            state.useravadar=data.useravadar
        }
    },
    actions: {
        updateSelectedTabIndex({ commit }, index) {
            commit('setSelectedTabIndex', index);
        },
        setSearchContent({ commit }, content){
            commit('setSearchContent',content);
            console.log(content)
        },
    },
    getters: {
        getSelectedTabIndex(state) {
            return state.selectedTabIndex;
        },
        getSearchContent(state){
            return state.searchContent;
        },
    }
});
