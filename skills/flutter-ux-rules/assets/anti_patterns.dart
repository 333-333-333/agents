// ❌ Redundant — adds zero value
/// The user's name.
final String userName;

// ❌ Restating the type signature
/// Returns a Future of User.
Future<User> getUser();

// ❌ Documenting generated code
/// From JSON factory.
factory User.fromJson(Map<String, dynamic> json);
