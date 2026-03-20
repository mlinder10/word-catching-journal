// import MockAPI from "./mock-api";

import API from "./api";

export * from "./responses";

export default new API(import.meta.env.VITE_SERVER_URL);
// export default new MockAPI();
