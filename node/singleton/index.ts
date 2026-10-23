import RegularCache = require("./cache/RegularCache");
import SingletonCache = require("./cache/SingletonCache");
import CacheService = require("./cacheServiceLocator/CacheService");
import ICache = require("./cacheServiceLocator/CacheService");
import ServiceLocator = require("./cacheServiceLocator/ServiceLocator");

const regularCache1 = new RegularCache();
const regularCache2 = new RegularCache();

regularCache1.set("user", "João");

console.log(regularCache1.get("user")); // João
console.log(regularCache2.get("user")); // undefined

const singletonCache1 = SingletonCache.getInstance();
const singletonCache2 = SingletonCache.getInstance();

singletonCache1.set("user", "João");

console.log(singletonCache1.get("user")); // João
console.log(singletonCache2.get("user")); // João

const serviceLocator = ServiceLocator.getInstance();
serviceLocator.register<ICache>("cacheService", new CacheService);

const serviceLocatorCache1 = serviceLocator.get<ICache>("cacheService");
const serviceLocatorCache2 = serviceLocator.get<ICache>("cacheService");

serviceLocatorCache1.set("user", "João");

console.log(serviceLocatorCache1.get("user"));
console.log(serviceLocatorCache2.get("user"));
