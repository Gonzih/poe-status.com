(ns poe.core)

(enable-console-print!)

(prn "hello world!")

(->
  (js/fetch "/data.json")
  (.then (fn [response]
           (if (.-ok response)
             (.json response)
             (throw (js/Error. "Something went wrong")))))
  (.then (fn [text]
           (js/console.log text))))
