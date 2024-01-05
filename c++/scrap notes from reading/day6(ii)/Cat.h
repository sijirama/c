
#include <string>
class Cat {
    public:
        Cat();
        Cat( std::string name);
        ~Cat();
        void Meow() const;
        void Purr() const;
        std::string getName();

    private:
        std::string name;
};
