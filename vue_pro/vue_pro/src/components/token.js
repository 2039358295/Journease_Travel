//存
export function setToken(token){
     return localStorage.setItem('token',token);
     //return sessionStorage.setItem('token',token);
    }
    //取
    export function getToken(){
     //为空时返回null
     return localStorage.getItem('token');
     //return sessionStorage.getItem('token');
    }
    //删
    export function removeToken(){
     return localStorage.removeItem('token');
     //return sessionStorage.removeItem('token');
    }