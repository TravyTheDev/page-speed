type Pet = {
  name: string;
  animal: string;
  favoriteFood: string;
};

type UserWithPet = {
  id: number;
  username: string;
  pet: Pet;
};

export async function initUserSelector() {
  const userSelect = document.getElementById('user-select') as HTMLSelectElement;
  const userName = document.getElementById('user-name')!;
  const pet = document.getElementById('pet')!;
  const animal = document.getElementById('animal')!;
  const favoriteFood = document.getElementById('favorite-food')!;

  const res = await fetch('http://localhost:8000/api/v1/get_users_with_pets_bad');
  const users: UserWithPet[] = await res.json();

  users.forEach((user) => {
    const option = document.createElement('option');
    option.value = String(user.id);
    option.textContent = user.username;
    userSelect.appendChild(option);
  });

  userSelect.addEventListener('change', (e) => {
    const selectedId = Number((e.target as HTMLSelectElement).value);
    const selectedUser = users.find((u) => u.id === selectedId);
    if (!selectedUser) return;

    userName.innerText = selectedUser.username;
    pet.innerText = selectedUser.pet.name;
    animal.innerText = selectedUser.pet.animal;
    favoriteFood.innerText = selectedUser.pet.favoriteFood;
  });
}
