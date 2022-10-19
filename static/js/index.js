const todos = document.querySelectorAll('li');

Array.from(todos).map(todo => {
    todo.addEventListener('click', () => {
        todo.classList.toggle('done');
    })
});